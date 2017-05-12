package registry

import (
	"testing"

	"github.com/liszt-code/liszt/pkg/registry/mocks"
	"github.com/stretchr/testify/assert"
)

type testHandlerObject struct {
	reg *mocks.Registrar
	svc *Service
}

func newTestHandlerObject() (tho *testHandlerObject) {
	m := new(mocks.Registrar)
	return &testHandlerObject{
		reg: m,
		svc: &Service{
			Registrar: m,
		},
	}
}

func (tho *testHandlerObject) teardown(t *testing.T) {
	assert.True(t, tho.reg.AssertExpectations(t))
}

func TestRegistryService(t *testing.T) {
	t.Run("ListUnitResidentsHandler", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
		})
		t.Run("unit not found", func(t *testing.T) {
		})
		t.Run("missing unit parameter", func(t *testing.T) {
		})
	})
}
