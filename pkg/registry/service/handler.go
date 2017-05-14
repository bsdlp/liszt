package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/bsdlp/apiutils"
	"github.com/liszt-code/liszt/pkg/registry"
)

// Service implements Registrar
type Service struct {
	Registrar registry.Registrar
	Logger    logrus.FieldLogger
}

// Router returns an http.Handler
func (svc *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/units":
		svc.GetUnitByNameHandler(w, r)
	case "/units/residents":
		svc.ListUnitResidentsHandler(w, r)
	case "/residents/register":
		svc.RegisterResidentHandler(w, r)
	case "/residents/move":
		svc.MoveResidentHandler(w, r)
	case "/residents/deregister":
		svc.DeregisterResidentHandler(w, r)
	default:
		apiutils.WriteError(w, apiutils.ErrNotFound)
	}
}

// ListUnitResidentsHandler implements registrar
func (svc *Service) ListUnitResidentsHandler(w http.ResponseWriter, r *http.Request) {
	unitID, err := strconv.ParseInt(r.URL.Query().Get("unit_id"), 10, 64)
	if err != nil {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, "unit_id required and must be an int64"))
		return
	}
	residents, err := svc.Registrar.ListUnitResidents(r.Context(), unitID)
	if err != nil {
		apiutils.WriteError(w, err)
		return
	}
	if len(residents) == 0 {
		apiutils.WriteError(w, apiutils.ErrNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(residents)
	return
}

// GetUnitByNameHandler implements registrar
func (svc *Service) GetUnitByNameHandler(w http.ResponseWriter, r *http.Request) {
	unit, err := svc.Registrar.GetUnitByName(r.Context(), r.URL.Query().Get("unit"))
	if err != nil {
		apiutils.WriteError(w, err)
		return
	}
	if unit == nil {
		apiutils.WriteError(w, apiutils.ErrNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(unit)
	return
}

// RegisterResidentHandler implements registrar
func (svc *Service) RegisterResidentHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		closeErr := r.Body.Close()
		if closeErr != nil {
			svc.Logger.Warnf("unable to close body: %s", closeErr.Error())
		}
	}()

	resident := new(registry.Resident)
	err := json.NewDecoder(r.Body).Decode(resident)
	if err != nil {
		apiutils.WriteError(w, apiutils.ErrBadRequest)
		return
	}

	registered, err := svc.Registrar.RegisterResident(r.Context(), resident)
	if err != nil {
		apiutils.WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = apiutils.WriteJSON(w, registered)
	if err != nil {
		svc.Logger.Warnf("error writing json: %s", err.Error())
	}
	return
}

// MoveResidentHandler implements registrar
func (svc *Service) MoveResidentHandler(w http.ResponseWriter, r *http.Request) {
	residentID, err := strconv.ParseInt(r.URL.Query().Get("resident_id"), 10, 64)
	if err != nil {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, "resident_id (int) is required"))
		return
	}

	unitID, err := strconv.ParseInt(r.URL.Query().Get("unit_id"), 10, 64)
	if err != nil {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, "unit_id (int) is required"))
		return
	}

	err = svc.Registrar.MoveResident(r.Context(), residentID, unitID)
	if err != nil {
		apiutils.WriteError(w, err)
		return
	}
	return
}

// DeregisterResidentHandler implements registrar
func (svc *Service) DeregisterResidentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		apiutils.WriteError(w, apiutils.ErrMethodNotAllowed)
		return
	}

	residentID, err := strconv.ParseInt(r.URL.Query().Get("resident_id"), 10, 64)
	if err != nil {
		apiutils.WriteError(w, apiutils.NewError(http.StatusBadRequest, "resident_id (int) is required"))
		return
	}

	err = svc.Registrar.DeregisterResident(r.Context(), residentID)
	if err != nil {
		apiutils.WriteError(w, err)
		return
	}
	return
}
