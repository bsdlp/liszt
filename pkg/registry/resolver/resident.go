package resolver

import (
	"github.com/liszt-code/liszt/pkg/registry"
	graphql "github.com/neelance/graphql-go"
	"github.com/sirupsen/logrus"
)

// ResidentResolver resolves residents
type ResidentResolver struct {
	resident  *registry.Resident
	ur        *UnitResolver
	registrar registry.Registrar
	logger    logrus.FieldLogger
}

// ID implements resident
func (rr *ResidentResolver) ID() graphql.ID {
	return graphql.ID(rr.resident.ID)
}

// Name implements resident
func (rr *ResidentResolver) Name() string {
	return rr.resident.String()
}

// Unit implements resident
func (rr *ResidentResolver) Unit() *UnitResolver {
	return rr.ur
}
