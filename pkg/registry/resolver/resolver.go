package resolver

import (
	"github.com/liszt-code/liszt/pkg/registry"
	"github.com/sirupsen/logrus"
)

// Resolver implements graphql resolvers
type Resolver struct {
	Registrar registry.Registrar
	Logger    logrus.FieldLogger
}
