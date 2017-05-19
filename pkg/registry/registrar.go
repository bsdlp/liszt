package registry

import "context"

// Registrar maintains a registry of units and residents
type Registrar interface {
	GetBuildingByID(ctx context.Context, buildingID string) (building *Building, err error)

	RegisterBuilding(ctx context.Context, in *Building) (building *Building, err error)

	DeregisterBuilding(ctx context.Context, buildingID string) (err error)

	ListBuildingUnits(ctx context.Context, buildingID string) (units []*Unit, err error)

	// register unit
	RegisterUnit(ctx context.Context, buildingID string, in *Unit) (unit *Unit, err error)

	DeregisterUnit(ctx context.Context, unitID string) (err error)

	ListUnitResidents(ctx context.Context, unitID int64) (residents []*Resident, err error)

	GetResidentByID(ctx context.Context, residentID string) (resident *Resident, err error)

	// adds a resident into the registry, optionally attaching the resident to
	// a unit if unitID is not empty.
	RegisterResident(ctx context.Context, resident *Resident) (returned *Resident, err error)

	DeregisterResident(ctx context.Context, residentID string) (err error)

	// moves a resident to a new unit
	MoveResidentIn(ctx context.Context, residentID, newUnitID string) (err error)

	// moves a resident out of a unitj
	MoveResidentOut(ctx context.Context, residentID, unitID string) (err error)
}

// Building describes a building
type Building struct {
	ID   string `dynamodbav:"building_id"`
	Name string
}

// Resident represents a resident in liszt
type Resident struct {
	ID string `dynamodbav:"resident_id"`

	Firstname  string
	Middlename string
	Lastname   string
}

func (res *Resident) String() string {
	return res.Lastname + ", " + res.Firstname + " " + res.Middlename
}

// Unit describes a unit in liszt
type Unit struct {
	ID   string
	Name string
}
