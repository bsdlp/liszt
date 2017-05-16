package rdbms

// Config holds configuration options for registry
type Config struct {
	DriverName     string `default:"mysql"`
	DataSourceName string `required:"true"`
}
