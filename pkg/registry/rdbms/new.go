package rdbms

import "github.com/jmoiron/sqlx"

// New returns a registry service
func New(cfg *Config) (reg *MySQLRegistrar, err error) {
	db, err := sqlx.Open(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		return
	}

	reg = &MySQLRegistrar{
		DB: db,
	}
	return
}
