package registry

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationBatchGetResidents(t *testing.T) {
	registeredResidents := make([]*Resident, 4)
	for i := range registeredResidents {
		req := &Resident{
			Firstname:  "Josiah",
			Middlename: "Edward",
			Lastname:   "Bartlet",
		}
		var err error
		registeredResidents[i], err = testRegistrar.RegisterResident(context.Background(), req)
		assert.NoError(t, err)
	}

	defer func() {
		for _, v := range registeredResidents {
			err := testRegistrar.DeregisterResident(context.Background(), v.ID)
			assert.NoError(t, err)
		}
	}()

	t.Run("batch get success", func(t *testing.T) {
		residentIDs := make([]string, len(registeredResidents))
		for i, v := range registeredResidents {
			residentIDs[i] = v.ID
		}

		residents, err := testRegistrar.batchGetResidents(context.Background(), residentIDs)
		assert.NoError(t, err)

		for _, v := range registeredResidents {
			assert.Contains(t, residents, v)
		}
	})

	err := testRegistrar.DeregisterResident(context.Background(), registeredResidents[3].ID)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("batch get some missing", func(t *testing.T) {
		residentIDs := make([]string, len(registeredResidents))
		for i, v := range registeredResidents {
			residentIDs[i] = v.ID
		}

		residents, err := testRegistrar.batchGetResidents(context.Background(), residentIDs)
		assert.NoError(t, err)

		assert.Len(t, residents, 3)
		for _, v := range registeredResidents[:2] {
			assert.Contains(t, residents, v)
		}

		assert.NotContains(t, residents, registeredResidents[3])
	})
}

func TestIntegrationResidents(t *testing.T) {
	t.Run("get nonexistent resident", func(t *testing.T) {
		assert := assert.New(t)
		resident, err := testRegistrar.GetResidentByID(context.Background(), "nonexistent")
		assert.NoError(err)
		assert.Nil(resident)
	})

	var registeredResident *Resident
	t.Run("register resident", func(t *testing.T) {
		assert := assert.New(t)
		req := &Resident{
			Firstname:  "Josiah",
			Middlename: "Edward",
			Lastname:   "Bartlet",
		}
		var err error
		registeredResident, err = testRegistrar.RegisterResident(context.Background(), req)
		assert.NoError(err)
		assert.NotNil(registeredResident)
		assert.NotEmpty(registeredResident.ID)
	})

	t.Run("get existing resident", func(t *testing.T) {
		assert := assert.New(t)
		resident, err := testRegistrar.GetResidentByID(context.Background(), registeredResident.ID)
		assert.NoError(err)
		assert.Equal(registeredResident, resident)
	})

	t.Run("deregister resident", func(t *testing.T) {
		assert := assert.New(t)
		err := testRegistrar.DeregisterResident(context.Background(), registeredResident.ID)
		assert.NoError(err)
	})

	t.Run("get deregistered resident", func(t *testing.T) {
		assert := assert.New(t)
		resident, err := testRegistrar.GetResidentByID(context.Background(), registeredResident.ID)
		assert.NoError(err)
		assert.Nil(resident)
	})

	t.Run("deregister nonexistent resident should not error", func(t *testing.T) {
		assert := assert.New(t)
		err := testRegistrar.DeregisterResident(context.Background(), "nonexistent")
		assert.NoError(err)
	})
}
