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
	mux.Post("/units/register", svc.RegisterUnit)
	mux.Post("/units/deregister", svc.DeregisterUnit)
	mux.Post("/residents/register", svc.RegisterResident)
	mux.Post("/residents/deregister", svc.DeregisterResident)
	mux.Post("/residents/move_in", svc.MoveResidentIn)
	mux.Post("/residents/move_out", svc.MoveResidentOut)
	return
}

type apiserver struct {
	registrar registry.Registrar
}
