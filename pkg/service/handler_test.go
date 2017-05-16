package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/liszt-code/liszt/pkg/registry/mocks"
	"github.com/stretchr/testify/assert"
)

type handlerTestObject struct {
	reg    *mocks.Registrar
	svc    *Service
	server *httptest.Server
}

func newHandlerTestObject() (hto *handlerTestObject) {
	reg := new(mocks.Registrar)
	svc := &Service{Registrar: reg}
	hto = &handlerTestObject{
		svc:    svc,
		reg:    reg,
		server: httptest.NewServer(svc),
	}
	return
}

func (hto *handlerTestObject) teardown(t *testing.T) {
	hto.server.Close()
	assert.True(t, hto.reg.AssertExpectations(t))
}

func TestHandler(t *testing.T) {
	t.Run("ListUnitResidentsHandler", func(t *testing.T) {
		t.Run("missing unit_id param", func(t *testing.T) {
			assert := assert.New(t)
			hto := newHandlerTestObject()
			defer hto.teardown(t)
			resp, err := http.Get(hto.server.URL + "/units/residents")
			assert.NoError(err)
			defer func() {
				closeErr := resp.Body.Close()
				assert.NoError(closeErr)
			}()
			assert.Equal(http.StatusBadRequest, resp.StatusCode)
		})
		t.Run("unit_id not an int64 type", func(t *testing.T) {
			assert := assert.New(t)
			hto := newHandlerTestObject()
			defer hto.teardown(t)
			resp, err := http.Get(hto.server.URL + "/units/residents?unit_id=true")
			assert.NoError(err)
			defer func() {
				closeErr := resp.Body.Close()
				assert.NoError(closeErr)
			}()
			assert.Equal(http.StatusBadRequest, resp.StatusCode)
		})
	})
}
