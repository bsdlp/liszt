package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bsdlp/apiutils"
	"github.com/liszt-code/liszt/pkg/registry"
)

// RegisterBuilding registers a building
func (svc *apiserver) RegisterBuilding(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			// do nothing
		}
	}()

	input := new(registry.Building)
	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	output, err := svc.registrar.RegisterBuilding(r.Context(), input)
	if err != nil {
		log.Println(err)
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

// DeregisterBuilding deregisters a building
func (svc *apiserver) DeregisterBuilding(w http.ResponseWriter, r *http.Request) {
	buildingID := r.URL.Query().Get("building_id")
	if buildingID == "" {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, "building_id is required"))
		return
	}

	err := svc.registrar.DeregisterBuilding(r.Context(), buildingID)
	if err != nil {
		apiutils.WriteError(w, err)
		return
	}
	return
}
