package rdbms

import (
	"context"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/liszt-code/liszt/pkg/registry"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// MySQLRegistrar implements registrar using mysql
type MySQLRegistrar struct {
	DB *sqlx.DB
}

const listUnitResidentsQuery = `select residents.* from residents
inner join units_residents
on residents.id = units_residents.resident
where units_residents.unit = ?;`

// ListUnitResidents implements registrar
func (reg *MySQLRegistrar) ListUnitResidents(ctx context.Context, unitID int64) (residents []*registry.Resident, err error) {
	rows, err := reg.DB.QueryxContext(ctx, listUnitResidentsQuery, unitID)
	if err != nil {
		return
	}
	defer func() {
		closeErr := rows.Close()
		if err == nil {
			err = closeErr
		}
	}()

	residents = []*registry.Resident{}
	for rows.Next() {
		resident := new(registry.Resident)
		err = rows.StructScan(resident)
		if err != nil {
			return
		}
		residents = append(residents, resident)
	}
	err = rows.Err()
	return
}

const getUnitByNameQuery = `select * from units
where units.name = ?;`

// GetUnitByName implmements registrar
func (reg *MySQLRegistrar) GetUnitByName(ctx context.Context, name string) (unit *registry.Unit, err error) {
	unit = new(registry.Unit)
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
func (reg *MySQLRegistrar) RegisterUnit(ctx context.Context, in *registry.Unit) (unit *registry.Unit, err error) {
	result, err := reg.DB.NamedExecContext(ctx, registerUnitQuery, in)
	if err != nil {
		return
	}
	unitID, err := result.LastInsertId()
	if err != nil {
		return
	}

	unit = new(registry.Unit)
	*unit = *in
	unit.ID = unitID
	return
}

const registerResidentQuery = `insert into residents (firstname, middlename, lastname)
values (:firstname, :middlename, :lastname);`

// RegisterResident implements registrar
func (reg *MySQLRegistrar) RegisterResident(ctx context.Context, resident *registry.Resident) (returned *registry.Resident, err error) {
	result, err := reg.DB.NamedExecContext(ctx, registerResidentQuery, resident)
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	returned = new(registry.Resident)
	*returned = *resident
	returned.ID = id
	return
}

const (
	moveOutResidentQuery = `delete from units_residents
	where units_residents.unit = ? and
	units_residents.resident = ?;`
	moveInResidentQuery = `insert into units_residents (unit, resident) values (?, ?);`
)

const (
	moveResidentForeignKeyConstraintError = "a foreign key constraint fails"
	moveResidentDuplicateEntry            = "Duplicate entry"
)

// MoveResidentIn implements registrar
func (reg *MySQLRegistrar) MoveResidentIn(ctx context.Context, residentID, unitID int64) (err error) {
	_, err = reg.DB.ExecContext(ctx, moveInResidentQuery, unitID, residentID)
	if err != nil {
		if strings.Contains(err.Error(), moveResidentForeignKeyConstraintError) {
			err = registry.ErrMissingUnitOrResident
			return
		}
		if strings.Contains(err.Error(), moveResidentDuplicateEntry) {
			err = registry.ErrResidentAlreadyInUnit
			return
		}
		return
	}
	return
}

// MoveResidentOut implements registrar
func (reg *MySQLRegistrar) MoveResidentOut(ctx context.Context, residentID, unitID int64) (err error) {
	result, err := reg.DB.ExecContext(ctx, moveOutResidentQuery, unitID, residentID)
	if err != nil {
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected == 0 {
		err = registry.ErrCannotMoveResidentOut
		return
	}
	return
}

const (
	deregisterResidentQuery = `delete from residents where residents.id = ?;`
)

// DeregisterResident implements registrar
func (reg *MySQLRegistrar) DeregisterResident(ctx context.Context, residentID int64) (err error) {
	result, err := reg.DB.ExecContext(ctx, deregisterResidentQuery, residentID)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected == 0 {
		err = registry.ErrResidentNotFound
		return
	}
	return
}
