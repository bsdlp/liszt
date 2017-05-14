package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20170513174841, Down20170513174841)
}

// Up20170513174841 updates the database to the new requirements
func Up20170513174841(tx *sql.Tx) (err error) {
	_, err = tx.Exec(`alter table units_residents
drop primary key,
drop column id,
add primary key(unit, resident)`)
	return
}

// Down20170513174841 should send the database back to the state it was from before Up was ran
func Down20170513174841(tx *sql.Tx) (err error) {
	_, err = tx.Exec(`alter table units_residents
drop primary key,
add column id int not null auto_increment,
add primary key(id)`)
	return
}
