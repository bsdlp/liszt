package internal

import (
	"encoding/json"
	"net/http"

	"github.com/bsdlp/apiutils"
	"github.com/liszt-code/liszt/pkg/registry"
)

// RegisterResident registers a resident
func (svc *apiserver) RegisterResident(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			// do nothing
		}
	}()

	input := new(registry.Resident)
	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	output, err := svc.registrar.RegisterResident(r.Context(), input)
	if err != nil {
		apiutils.WriteError(w, err)
		return
	}

	err = apiutils.WriteJSON(w, output)
	if err != nil {
		apiutils.WriteError(w, err)
		return
	}
	return
}

// DeregisterResident deregisters a resident
func (svc *apiserver) DeregisterResident(w http.ResponseWriter, r *http.Request) {
	residentID := r.URL.Query().Get("resident_id")
	if residentID == "" {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, "resident_id is required"))
		return
	}

	err := svc.registrar.DeregisterResident(r.Context(), residentID)
	if err != nil {
		apiutils.WriteError(w, err)
		return
	}
	return
}

// MoveResidentIn moves a resident into a unit
func (svc *apiserver) MoveResidentIn(w http.ResponseWriter, r *http.Request) {
	residentID := r.URL.Query().Get("resident_id")
	if residentID == "" {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, "resident_id is required"))
		return
	}

	unitID := r.URL.Query().Get("unit_id")
	if unitID == "" {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, "unit_id is required"))
		return
	}

	err := svc.registrar.MoveResidentIn(r.Context(), residentID, unitID)
	if err != nil {
		apiutils.WriteError(w, err)
		return
	}
	return
}

// MoveResidentOut moves a resident out of a unit
func (svc *apiserver) MoveResidentOut(w http.ResponseWriter, r *http.Request) {
	residentID := r.URL.Query().Get("resident_id")
	if residentID == "" {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, "resident_id is required"))
		return
	}

	unitID := r.URL.Query().Get("unit_id")
	if unitID == "" {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, "unit_id is required"))
		return
	}

	err := svc.registrar.MoveResidentOut(r.Context(), residentID, unitID)
	if err != nil {
		apiutils.WriteError(w, err)
		return
	}
	return
}
