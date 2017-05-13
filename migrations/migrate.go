package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

// Migrate migrates
func Migrate(db *sql.DB) (err error) {
	err = goose.SetDialect("mysql")
	if err != nil {
		return
	}

	err = goose.Up(db, "")
	if err != nil {
		return
	}
	return
}
