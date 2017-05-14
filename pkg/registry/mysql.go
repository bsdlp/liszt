package registry

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/bsdlp/apiutils"
	"github.com/jmoiron/sqlx"

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
func (reg *MySQLRegistrar) ListUnitResidents(ctx context.Context, unitID int64) (residents []*Resident, err error) {
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

	residents = []*Resident{}
	for rows.Next() {
		resident := new(Resident)
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
func (reg *MySQLRegistrar) RegisterResident(ctx context.Context, resident *Resident) (returned *Resident, err error) {
	result, err := reg.DB.NamedExecContext(ctx, registerResidentQuery, resident)
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	returned = new(Resident)
	*returned = *resident
	returned.ID = id
	return
}

const (
	moveOutResidentQuery = `delete from units_residents where units_residents.resident = ?;`
	moveInResidentQuery  = `insert into units_residents (unit, resident) values (?, ?);`
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

// ErrResidentNotFound is returned when resident is not found
var ErrResidentNotFound = apiutils.NewError(http.StatusNotFound, "resident not found")

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
		err = ErrResidentNotFound
		return
	}
	return
}
