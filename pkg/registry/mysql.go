package registry

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// MySQLRegistrar implements registrar using mysql
type MySQLRegistrar struct {
	DB *sqlx.DB
}

const listUnitResidentsQuery = `select * from residents
inner join units_residents
on residents.id = units_residents.resident
where residents.id is ?;`

// ListUnitResidents implements registrar
func (reg *MySQLRegistrar) ListUnitResidents(ctx context.Context, unitID int64) (residents []*Resident, err error) {
	err = reg.DB.SelectContext(ctx, &residents, listUnitResidentsQuery, unitID)
	return
}

const getUnitByNameQuery = `select * from units
where units.name = ?;`

// GetUnitByName implmements registrar
func (reg *MySQLRegistrar) GetUnitByName(ctx context.Context, name string) (unit *Unit, err error) {
	unit = new(Unit)
	err = reg.DB.GetContext(ctx, unit, getUnitByNameQuery, name)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			unit = nil
		}
		return
	}
	return
}

const registerUnitQuery = `insert into units (name)
values (:name);`

// RegisterUnit registers a unit
func (reg *MySQLRegistrar) RegisterUnit(ctx context.Context, in *Unit) (unit *Unit, err error) {
	result, err := reg.DB.NamedExecContext(ctx, registerUnitQuery, in)
	if err != nil {
		return
	}
	unitID, err := result.LastInsertId()
	if err != nil {
		return
	}

	unit = new(Unit)
	*unit = *in
	unit.ID = unitID
	return
}

const registerResidentQuery = `insert into residents (firstname, middlename, lastname)
values (:firstname, :middlename, :lastname);`

// RegisterResident implements registrar
func (reg *MySQLRegistrar) RegisterResident(ctx context.Context, resident *Resident) (err error) {
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
func (reg *MySQLRegistrar) MoveResident(ctx context.Context, residentID, newUnitID int64) (err error) {
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
func (reg *MySQLRegistrar) DeregisterResident(ctx context.Context, residentID int64) (err error) {
	_, err = reg.DB.ExecContext(ctx, moveOutResidentQuery, residentID)
	return
}
