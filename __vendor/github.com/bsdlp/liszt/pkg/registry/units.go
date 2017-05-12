package registry

// Unit describes a unit in liszt
type Unit struct {
	ID   int64
	Name string `db:"name"`
}
