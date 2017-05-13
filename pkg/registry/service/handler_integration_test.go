package service

import (
	"context"
	"encoding/json"
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

func TestIntegrationHandler(t *testing.T) {
	hito := newHandlerIntegrationTestObject(t)
	defer hito.teardown(t)

	t.Run("GetUnitByNameHandler", func(t *testing.T) {
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
	})
}
