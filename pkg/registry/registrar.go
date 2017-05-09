package registry

import "context"

// Registrar maintains a registry of units and residents
type Registrar interface {
	GetUnitResidents(ctx context.Context, unitID string) (residents []*Resident, err error)

	// get unit by unit name
	GetUnitByName(ctx context.Context, name string) (unit *Unit, err error)

	// adds a resident into the registry, optionally attaching the resident to
	// a unit if unitID is not empty.
	RegisterResident(ctx context.Context, resident *Resident, unitID string) (err error)

	// moves a resident to a new unit
	MoveResident(ctx context.Context, residentID, newUnitID string) (err error)

	// removes a user from the directory entirely
	DeregisterResident(ctx context.Context, residentID string) (err error)
}
