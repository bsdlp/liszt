package registry

// Resident represents a resident in liszt
type Resident struct {
	ID string

	Firstname  string
	Middlename string
	Lastname   string
}

func (res *Resident) String() string {
	return res.Lastname + ", " + res.Firstname + " " + res.Middlename
}
