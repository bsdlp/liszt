package internal

import (
	"encoding/json"
	"net/http"

	"github.com/bsdlp/apiutils"
	"github.com/liszt-code/liszt/pkg/registry"
)

// RegisterUnit registers a unit
func (svc *apiserver) RegisterUnit(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			// do nothing
		}
	}()

	input := new(registry.Unit)
	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	buildingID := r.URL.Query().Get("building_id")
	if buildingID == "" {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, "building_id is required"))
		return
	}

	output, err := svc.registrar.RegisterUnit(r.Context(), buildingID, input)
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

// DeregisterUnit deregisters a unit
func (svc *apiserver) DeregisterUnit(w http.ResponseWriter, r *http.Request) {
	unitID := r.URL.Query().Get("unit_id")
	if unitID == "" {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, "unit_id is required"))
		return
	}

	err := svc.registrar.DeregisterUnit(r.Context(), unitID)
	if err != nil {
		apiutils.WriteError(w, err)
		return
	}
	return
}
