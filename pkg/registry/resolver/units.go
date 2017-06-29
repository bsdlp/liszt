package resolver

import (
	"context"

	"github.com/liszt-code/liszt/pkg/registry"
	graphql "github.com/neelance/graphql-go"
	"github.com/sirupsen/logrus"
)

// UnitResolver implements unit resolver
type UnitResolver struct {
	br        *BuildingResolver
	unit      *registry.Unit
	registrar registry.Registrar
	logger    logrus.FieldLogger
}

// ID implements unit
func (ur *UnitResolver) ID() graphql.ID {
	return graphql.ID(ur.unit.ID)
}

// Name implements unit
func (ur *UnitResolver) Name() string {
	return ur.unit.Name
}

// Building implements unit
func (ur *UnitResolver) Building() *BuildingResolver {
	return ur.br
}

// Residents implements units
func (ur *UnitResolver) Residents() (rr []*ResidentResolver) {
	residents, err := ur.registrar.ListUnitResidents(context.TODO(), ur.unit.ID)
	if err != nil {
		ur.logger.Error(err)
		return
	}

	rr = make([]*ResidentResolver, len(residents))
	for i, resident := range residents {
		rr[i] = &ResidentResolver{
			resident:  resident,
			ur:        ur,
			registrar: ur.registrar,
			logger:    ur.logger,
		}
	}
	return
}
