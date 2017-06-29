package resolver

import (
	"github.com/liszt-code/liszt/pkg/registry"
	"github.com/sirupsen/logrus"
)

// Resolver implements graphql resolvers
type Resolver struct {
	registrar registry.Registrar
	logger    logrus.FieldLogger
}
