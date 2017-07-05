package internal

import (
	"github.com/go-chi/chi"
	"github.com/liszt-code/liszt/pkg/registry"
)

// NewCRUDService returns a CRUD apiserver
func NewCRUDService(registrar registry.Registrar) (mux *chi.Mux) {
	svc := &apiserver{registrar: registrar}
	mux = chi.NewMux()
	mux.Post("/buildings/register", svc.RegisterBuilding)
	mux.Post("/buildings/deregister", svc.DeregisterBuilding)
	return
}

type apiserver struct {
	registrar registry.Registrar
}
