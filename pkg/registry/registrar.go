package registry

import (
	"context"
	"net/http"

	"github.com/bsdlp/apiutils"
)

// Registrar maintains a registry of units and residents
type Registrar interface {
	ListUnitResidents(ctx context.Context, unitID int64) (residents []*Resident, err error)

	// get unit by unit name
	GetUnitByName(ctx context.Context, name string) (unit *Unit, err error)

	// register unit
	RegisterUnit(ctx context.Context, in *Unit) (unit *Unit, err error)

	// adds a resident into the registry, optionally attaching the resident to
	// a unit if unitID is not empty.
	RegisterResident(ctx context.Context, resident *Resident) (returned *Resident, err error)

	// moves a resident to a new unit
	MoveResidentIn(ctx context.Context, residentID, newUnitID int64) (err error)

	// moves a resident out of a unitj
	MoveResidentOut(ctx context.Context, residentID, unitID int64) (err error)

	// removes a user from the directory entirely
	DeregisterResident(ctx context.Context, residentID int64) (err error)
}

var (
	// ErrMissingUnitOrResident is returned when trying to move a missing resident or to a missing unit
	ErrMissingUnitOrResident = apiutils.NewError(http.StatusUnprocessableEntity, "specified unit or resident does not exist")

	// ErrResidentAlreadyInUnit is returned when trying to move resident into a unit where the resident already resides
	ErrResidentAlreadyInUnit = apiutils.NewError(http.StatusUnprocessableEntity, "resident already resides in specified unit")

	// ErrCannotMoveResidentOut is returned when trying to move a resident out of a unit and it doesn't work.
	ErrCannotMoveResidentOut = apiutils.NewError(http.StatusUnprocessableEntity, "cannot move resident out, either the unit/resident does not exist or the resident does not reside in unit")

	// ErrResidentNotFound is returned when resident is not found
	ErrResidentNotFound = apiutils.NewError(http.StatusNotFound, "resident not found")
)
