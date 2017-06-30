package registry

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnits(t *testing.T) {
	t.Run("list units for a nonexistent building", func(t *testing.T) {
		assert := assert.New(t)
		units, err := testRegistrar.ListBuildingUnits(context.Background(), "nonexistent")
		assert.NoError(err)
		assert.Empty(units)
	})

	registeredBuilding, err := testRegistrar.RegisterBuilding(context.Background(), &Building{
		Name: getULID().String(),
	})
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := testRegistrar.DeregisterBuilding(context.Background(), registeredBuilding.ID)
		if err != nil {
			t.Fatal(err)
		}
	}()

	t.Run("list nonexistent units for an existing building", func(t *testing.T) {
		units, err := testRegistrar.ListBuildingUnits(context.Background(), registeredBuilding.ID)
		assert.NoError(t, err)
		assert.Empty(t, units)
	})

	registeredUnits := make([]*Unit, 4)
	t.Run("register multiple units to a building", func(t *testing.T) {
		for i := range registeredUnits {
			registeredUnit, err := testRegistrar.RegisterUnit(context.Background(), registeredBuilding.ID, &Unit{
				Name: getULID().String(),
			})
			assert.NoError(t, err)
			assert.NotEmpty(t, registeredUnit)
			registeredUnits[i] = registeredUnit
		}

		t.Run("list units for a registered building", func(t *testing.T) {
			units, err := testRegistrar.ListBuildingUnits(context.Background(), registeredBuilding.ID)
			assert.NoError(t, err)
			assert.Len(t, units, len(registeredUnits))
			for _, registeredUnit := range registeredUnits {
				assert.Contains(t, units, registeredUnit)
			}
		})

		t.Run("move multiple residents in", func(t *testing.T) {
			residents := make([]*Resident, 2)
			for i := range residents {
				resident, err := testRegistrar.RegisterResident(context.Background(), &Resident{
					Firstname:  "Josiah",
					Middlename: "Edward",
					Lastname:   "Bartlet",
				})
				if err != nil {
					t.Fatal(err)
				}

				err = testRegistrar.MoveResidentIn(context.Background(), resident.ID, registeredUnits[0].ID)
				assert.NoError(t, err)

				defer func() {
					err := testRegistrar.DeregisterResident(context.Background(), resident.ID)
					assert.NoError(t, err)
				}()
				residents[i] = resident
			}

			t.Run("list unit residents", func(t *testing.T) {
				movedInResidents, err := testRegistrar.ListUnitResidents(context.Background(), registeredUnits[0].ID)
				assert.NoError(t, err)
				assert.Len(t, movedInResidents, 2)
				for _, v := range movedInResidents {
					assert.Contains(t, residents, v)
				}
			})

			t.Run("move out one resident", func(t *testing.T) {
				err := testRegistrar.MoveResidentOut(context.Background(), residents[1].ID, registeredUnits[0].ID)
				assert.NoError(t, err)

				movedInResidents, err := testRegistrar.ListUnitResidents(context.Background(), registeredUnits[0].ID)
				assert.NoError(t, err)
				assert.Len(t, movedInResidents, 1)
				assert.Equal(t, residents[0], movedInResidents[0])
			})

			t.Run("move out last resident", func(t *testing.T) {
				err := testRegistrar.MoveResidentOut(context.Background(), residents[0].ID, registeredUnits[0].ID)
				assert.NoError(t, err)

				movedInResidents, err := testRegistrar.ListUnitResidents(context.Background(), registeredUnits[0].ID)
				assert.NoError(t, err)
				assert.Len(t, movedInResidents, 0)
			})
		})

		t.Run("deregister units", func(t *testing.T) {
			for _, unit := range registeredUnits {
				err := testRegistrar.DeregisterUnit(context.Background(), unit.ID)
				assert.NoError(t, err)
			}

			t.Run("units are deregistered", func(t *testing.T) {
				units, err := testRegistrar.ListBuildingUnits(context.Background(), registeredBuilding.ID)
				assert.NoError(t, err)
				assert.Empty(t, units)
			})
		})
	})
}
