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

const listUnitResidentsQuery = `select * from residents
inner join units_residents
on residents.id = units_residents.resident
where residents.id is ?;`

// ListUnitResidents implements registrar
func (reg *registrar) ListUnitResidents(ctx context.Context, unitID int64) (residents []*Resident, err error) {
	err = reg.DB.SelectContext(ctx, residents, listUnitResidentsQuery, unitID)
	return
}

const getUnitByNameQuery = `select * from units
where units.name = ?;`

// GetUnitByName implmements registrar
func (reg *registrar) GetUnitByName(ctx context.Context, name string) (unit *Unit, err error) {
	err = reg.DB.SelectContext(ctx, unit, getUnitByNameQuery, name)
	return
}

const registerResidentQuery = `insert into residents (firstname, middlename, lastname)
values (:firstname, :middlename, :lastname);`

// RegisterResident implements registrar
func (reg *registrar) RegisterResident(ctx context.Context, resident *Resident) (err error) {
	result, err := reg.DB.NamedExecContext(ctx, registerResidentQuery, resident)
	if err != nil {
		return
	}
	resident.ID, err = result.LastInsertId()
	return
}

const (
	moveOutResidentQuery = `delete from units_residents ur
where ur.resident = ?;`
	moveInResidentQuery = `insert into units_residents (unit, resident)
values (?, ?);`
)

// MoveResident implements registrar
func (reg *registrar) MoveResident(ctx context.Context, residentID, newUnitID int64) (err error) {
	_, err = reg.DB.ExecContext(ctx, moveOutResidentQuery, residentID)
	if err != nil {
		return
	}
	_, err = reg.DB.ExecContext(ctx, moveInResidentQuery, newUnitID, residentID)
	return
}

const (
	deregisterResidentQuery = `delete from residents where residents.id = ?;`
)

// DeregisterResident implements registrar
func (reg *registrar) DeregisterResident(ctx context.Context, residentID int64) (err error) {
	_, err = reg.DB.ExecContext(ctx, moveOutResidentQuery, residentID)
	return
}
