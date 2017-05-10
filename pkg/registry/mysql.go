package registry

import (
	"context"

	"github.com/jmoiron/sqlx"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

type registrar struct {
	DB *sqlx.DB
}

// GetUnitResidents implements registrar
func (req *registrar) GetUnitResidents(ctx context.Context, unitID string) (residents []*Resident, err error) {
	return
}

// GetUnitByName implmements registrar
func (req *registrar) GetUnitByName(ctx context.Context, name string) (unit *Unit, err error) {
	return
}

// RegisterResident implements registrar
func (req *registrar) RegisterResident(ctx context.Context, resident *Resident, unitID string) (err error) {
	return
}

// MoveResident implements registrar
func (req *registrar) MoveResident(ctx context.Context, residentID, newUnitID string) (err error) {
	return
}

// DeregisterResident implements registrar
func (req *registrar) DeregisterResident(ctx context.Context, residentID string) (err error) {
	return
}
