package registry

import (
	"context"

	"github.com/jmoiron/sqlx"
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
	RegisterResident(ctx context.Context, resident *Resident) (err error)

	// moves a resident to a new unit
	MoveResident(ctx context.Context, residentID, newUnitID int64) (err error)

	// removes a user from the directory entirely
	DeregisterResident(ctx context.Context, residentID int64) (err error)
}

// New returns a registry service
func New(cfg *Config) (svc *Service, err error) {
	db, err := sqlx.Open(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		return
	}

	svc = &Service{
		Registrar: &registrar{
			DB: db,
		},
	}
	return
}
