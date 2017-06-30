package resolver

import (
	"context"

	"github.com/liszt-code/liszt/pkg/registry"
	graphql "github.com/neelance/graphql-go"
	"github.com/sirupsen/logrus"
)

// Building retrieves a building by ID
func (r *Resolver) Building(args struct{ BuildingID graphql.ID }) (br *BuildingResolver) {
	building, err := r.Registrar.GetBuildingByID(context.TODO(), string(args.BuildingID))
	if err != nil {
		r.Logger.Error(err)
		return
	}

	br = &BuildingResolver{
		building:  building,
		registrar: r.Registrar,
		logger:    r.Logger,
	}
	return
}

// BuildingResolver implements building
type BuildingResolver struct {
	building  *registry.Building
	registrar registry.Registrar
	logger    logrus.FieldLogger
}

// ID implmements building
func (br *BuildingResolver) ID() graphql.ID {
	return graphql.ID(br.building.ID)
}

// Name implmements building
func (br *BuildingResolver) Name() string {
	return br.building.Name
}

// Address implements building
func (br *BuildingResolver) Address() string {
	return br.building.Address
}

// Units implements building
func (br *BuildingResolver) Units() (ur []*UnitResolver) {
	units, err := br.registrar.ListBuildingUnits(context.TODO(), br.building.ID)
	if err != nil {
		br.logger.Error(err)
		return
	}

	ur = make([]*UnitResolver, len(units))

	for i, unit := range units {
		ur[i] = &UnitResolver{
			br:        br,
			unit:      unit,
			registrar: br.registrar,
			logger:    br.logger,
		}
	}
	return
}

// Buildings returns buildings
func (r *Resolver) Buildings() (br []*BuildingResolver) {
	buildings, err := r.Registrar.ListBuildings(context.TODO())
	if err != nil {
		r.Logger.Error(err)
		return
	}

	br = make([]*BuildingResolver, len(buildings))
	for i, building := range buildings {
		br[i] = &BuildingResolver{building: building}
	}
	return
}
