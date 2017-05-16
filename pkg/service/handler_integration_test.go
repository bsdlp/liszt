package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/bsdlp/apiutils"
	"github.com/jmoiron/sqlx"
	"github.com/liszt-code/liszt/migrations"
	"github.com/liszt-code/liszt/pkg/registry"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

type handlerIntegrationTestObject struct {
	databaseName string
	db           *sqlx.DB
	svc          *Service
	server       *httptest.Server
}

func newHandlerIntegrationTestObject(t *testing.T) (hito *handlerIntegrationTestObject) {
	db, err := sqlx.Open("mysql", "root:@/")
	if err != nil {
		t.Fatal(err)
	}

	testDatabaseName := "test" + strconv.FormatInt(time.Now().Unix(), 10)

	_, err = db.Exec("create database " + testDatabaseName)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec("use " + testDatabaseName)
	if err != nil {
		t.Fatal(err)
	}

	err = migrations.Migrate(db.DB)
	if err != nil {
		t.Fatal(err)
	}

	svc := &Service{
		Registrar: &registry.MySQLRegistrar{
			DB: db,
		},
	}
	hito = &handlerIntegrationTestObject{
		databaseName: testDatabaseName,
		db:           db,
		svc:          svc,
		server:       httptest.NewServer(svc),
	}
	return
}

func (hito *handlerIntegrationTestObject) teardown(t *testing.T) {
	_, err := hito.db.Exec("drop database " + hito.databaseName)
	assert.NoError(t, err)
	assert.NoError(t, hito.db.Close())
	hito.server.Close()
}

func TestIntegrationListUnitResidentsHandler(t *testing.T) {
	hito := newHandlerIntegrationTestObject(t)
	defer hito.teardown(t)

	existingUnitName := uuid.NewV4().String()
	registeredUnit, err := hito.svc.Registrar.RegisterUnit(context.Background(), &registry.Unit{Name: existingUnitName})
	if err != nil {
		t.Fatal(err)
	}
	registeredUnitID := strconv.FormatInt(registeredUnit.ID, 10)

	expectedResidents := make([]*registry.Resident, 4)
	for i := range expectedResidents {
		expectedResidents[i], err = hito.svc.Registrar.RegisterResident(context.Background(), &registry.Resident{})
		if err != nil {
			t.Fatal(err)
		}
		err = hito.svc.Registrar.MoveResidentIn(context.Background(), expectedResidents[i].ID, registeredUnit.ID)
		if err != nil {
			t.Fatal(err)
		}
	}

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		resp, err := http.Get(hito.server.URL + "/units/residents?unit_id=" + registeredUnitID)
		assert.NoError(err)
		defer func() {
			assert.NoError(resp.Body.Close())
		}()
		assert.Equal(http.StatusOK, resp.StatusCode)

		var residents []*registry.Resident
		err = json.NewDecoder(resp.Body).Decode(&residents)
		assert.NoError(err)
		assert.Equal(expectedResidents, residents)
	})

	t.Run("unit not found", func(t *testing.T) {
		assert := assert.New(t)
		resp, err := http.Get(hito.server.URL + "/units/residents?unit_id=1234")
		assert.NoError(err)
		defer func() {
			assert.NoError(resp.Body.Close())
		}()
		assert.Equal(http.StatusNotFound, resp.StatusCode)

		var errObj apiutils.ErrorObject
		err = json.NewDecoder(resp.Body).Decode(&errObj)
		assert.NoError(err)
		assert.Equal(apiutils.ErrNotFound, errObj)
	})
}

func TestIntegrationGetUnitByNameHandler(t *testing.T) {
	hito := newHandlerIntegrationTestObject(t)
	defer hito.teardown(t)
	existingUnitName := uuid.NewV4().String()
	registeredUnit, err := hito.svc.Registrar.RegisterUnit(context.Background(), &registry.Unit{Name: existingUnitName})
	if err != nil {
		t.Fatal(err)
	}
	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		resp, err := http.Get(hito.server.URL + "/units?unit=" + existingUnitName)
		assert.NoError(err)
		assert.Equal(http.StatusOK, resp.StatusCode)
		defer func() {
			assert.NoError(resp.Body.Close())
		}()

		retrievedUnit := new(registry.Unit)
		err = json.NewDecoder(resp.Body).Decode(retrievedUnit)
		assert.NoError(err)
		assert.Equal(registeredUnit, retrievedUnit)
	})
	t.Run("unit not found", func(t *testing.T) {
		assert := assert.New(t)
		resp, err := http.Get(hito.server.URL + "/units?unit=" + uuid.NewV4().String())
		assert.NoError(err)
		assert.Equal(http.StatusNotFound, resp.StatusCode)
		defer func() {
			assert.NoError(resp.Body.Close())
		}()

		var errObj apiutils.ErrorObject
		err = json.NewDecoder(resp.Body).Decode(&errObj)
		assert.NoError(err)
		assert.Equal(apiutils.ErrNotFound, errObj)
	})
}

func TestIntegrationRegisterResidentHandler(t *testing.T) {
	assert := assert.New(t)
	hito := newHandlerIntegrationTestObject(t)
	defer hito.teardown(t)

	resident := &registry.Resident{
		Firstname:  "Josiah",
		Middlename: "Edward",
		Lastname:   "Bartlet",
	}

	var bs bytes.Buffer
	err := json.NewEncoder(&bs).Encode(resident)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.Post(hito.server.URL+"/residents/register", "application/json", &bs)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		closeErr := resp.Body.Close()
		assert.NoError(closeErr)
	}()
	registeredResident := new(registry.Resident)
	err = json.NewDecoder(resp.Body).Decode(registeredResident)
	assert.NoError(err)
	assert.NotEmpty(registeredResident)

	row := hito.db.QueryRowx("select * from residents where residents.id = ?", registeredResident.ID)
	storedResident := new(registry.Resident)
	err = row.StructScan(storedResident)
	if err != nil {
		assert.NoError(err)
	}
	assert.NoError(row.Err())
	assert.Equal(storedResident, registeredResident)
}

func TestIntegrationMoveResidentInHandler(t *testing.T) {
	hito := newHandlerIntegrationTestObject(t)
	defer hito.teardown(t)

	unit := &registry.Unit{
		Name: "testunit",
	}
	registeredUnit, err := hito.svc.Registrar.RegisterUnit(context.Background(), unit)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("nonexistent resident", func(t *testing.T) {
		assert := assert.New(t)
		resp, err := http.Post(hito.server.URL+fmt.Sprintf("/residents/move_in?resident_id=1&unit_id=%d", registeredUnit.ID), "application/json", nil)
		assert.NoError(err)
		defer func() {
			closeErr := resp.Body.Close()
			assert.NoError(closeErr)
		}()

		assert.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
		var errObj apiutils.ErrorObject
		err = json.NewDecoder(resp.Body).Decode(&errObj)
		assert.NoError(err)
		assert.Equal(registry.ErrMissingUnitOrResident, errObj)
	})

	resident := &registry.Resident{
		Firstname:  "josaiah",
		Middlename: "edward",
		Lastname:   "bartlet",
	}
	registeredResident, err := hito.svc.Registrar.RegisterResident(context.Background(), resident)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("nonexistent unit", func(t *testing.T) {
		assert := assert.New(t)
		resp, err := http.Post(hito.server.URL+fmt.Sprintf("/residents/move_in?resident_id=%d&unit_id=5", registeredResident.ID), "application/json", nil)
		assert.NoError(err)
		defer func() {
			assert.NoError(resp.Body.Close())
		}()

		assert.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
		var errObj apiutils.ErrorObject
		err = json.NewDecoder(resp.Body).Decode(&errObj)
		assert.NoError(err)
		assert.Equal(registry.ErrMissingUnitOrResident, errObj)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		resp, err := http.Post(hito.server.URL+fmt.Sprintf("/residents/move_in?resident_id=%d&unit_id=%d", registeredResident.ID, registeredUnit.ID), "application/json", nil)
		assert.NoError(err)
		defer func() {
			assert.NoError(resp.Body.Close())
		}()

		assert.Equal(http.StatusCreated, resp.StatusCode)
	})

	t.Run("moving resident into a unit where the resident already resides", func(t *testing.T) {
		assert := assert.New(t)
		resp, err := http.Post(hito.server.URL+fmt.Sprintf("/residents/move_in?resident_id=%d&unit_id=%d", registeredResident.ID, registeredUnit.ID), "application/json", nil)
		assert.NoError(err)
		defer func() {
			closeErr := resp.Body.Close()
			assert.NoError(closeErr)
		}()

		assert.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
		var errObj apiutils.ErrorObject
		err = json.NewDecoder(resp.Body).Decode(&errObj)
		assert.NoError(err)
		assert.Equal(registry.ErrResidentAlreadyInUnit, errObj)
	})
}

func TestIntegrationMoveResidentOutHandler(t *testing.T) {
	hito := newHandlerIntegrationTestObject(t)
	defer hito.teardown(t)
	unit := &registry.Unit{
		Name: "testunit",
	}
	registeredUnit, err := hito.svc.Registrar.RegisterUnit(context.Background(), unit)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("nonexistent resident", func(t *testing.T) {
		assert := assert.New(t)
		resp, err := http.Post(hito.server.URL+fmt.Sprintf("/residents/move_out?resident_id=1&unit_id=%d", registeredUnit.ID), "application/json", nil)
		assert.NoError(err)
		defer func() {
			closeErr := resp.Body.Close()
			assert.NoError(closeErr)
		}()

		assert.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
		var errObj apiutils.ErrorObject
		err = json.NewDecoder(resp.Body).Decode(&errObj)
		assert.NoError(err)
		assert.Equal(registry.ErrCannotMoveResidentOut, errObj)
	})

	resident := &registry.Resident{
		Firstname:  "josaiah",
		Middlename: "edward",
		Lastname:   "bartlet",
	}
	registeredResident, err := hito.svc.Registrar.RegisterResident(context.Background(), resident)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("nonexistent unit", func(t *testing.T) {
		assert := assert.New(t)
		resp, err := http.Post(hito.server.URL+fmt.Sprintf("/residents/move_out?resident_id=%d&unit_id=5", registeredResident.ID), "application/json", nil)
		assert.NoError(err)
		defer func() {
			assert.NoError(resp.Body.Close())
		}()

		assert.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
		var errObj apiutils.ErrorObject
		err = json.NewDecoder(resp.Body).Decode(&errObj)
		assert.NoError(err)
		assert.Equal(registry.ErrCannotMoveResidentOut, errObj)
	})

	err = hito.svc.Registrar.MoveResidentIn(context.Background(), registeredResident.ID, registeredUnit.ID)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		resp, err := http.Post(hito.server.URL+fmt.Sprintf("/residents/move_out?resident_id=%d&unit_id=%d", registeredResident.ID, registeredUnit.ID), "application/json", nil)
		assert.NoError(err)
		defer func() {
			assert.NoError(resp.Body.Close())
		}()

		assert.Equal(http.StatusOK, resp.StatusCode)
	})
}

func TestIntegrationDeregisterResidentHandler(t *testing.T) {
	hito := newHandlerIntegrationTestObject(t)
	defer hito.teardown(t)

	resident := &registry.Resident{
		Firstname:  "Josiah",
		Middlename: "Edward",
		Lastname:   "Bartlet",
	}
	registeredResident, err := hito.svc.Registrar.RegisterResident(context.Background(), resident)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("nonexistent resident", func(t *testing.T) {
		assert := assert.New(t)
		req, err := http.NewRequest(http.MethodDelete, hito.server.URL+"/residents/deregister?resident_id=14919414", nil)
		assert.NoError(err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(err)
		assert.Equal(http.StatusNotFound, resp.StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		req, err := http.NewRequest(http.MethodDelete, hito.server.URL+fmt.Sprintf("/residents/deregister?resident_id=%d", registeredResident.ID), nil)
		assert.NoError(err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(err)
		assert.Equal(http.StatusOK, resp.StatusCode)

		row := hito.db.QueryRowx("select * from residents where residents.id = ?", registeredResident.ID)
		assert.NoError(row.Err())
	})
}
