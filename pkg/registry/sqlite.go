package registry

import (
	"context"

	"github.com/jmoiron/sqlx"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type registrar struct {
	DB *sqlx.DB
}

// GetUnitResidents implements registrar
func (req *registrar) GetUnitResidents(ctx context.Context, unitID string) (residents []*Resident, err error)

// GetUnitByName implmements registrar
func (req *registrar) GetUnitByName(ctx context.Context, name string) (unit *Unit, err error)

// RegisterResident implements registrar
func (req *registrar) RegisterResident(ctx context.Context, resident *Resident, unitID string) (err error)

// MoveResident implements registrar
func (req *registrar) MoveResident(ctx context.Context, residentID, newUnitID string) (err error)

// DeregisterResident implements registrar
func (req *registrar) DeregisterResident(ctx context.Context, residentID string) (err error)
