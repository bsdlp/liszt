package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20170511195623, Down20170511195623)
}

// Up20170511195623 ups
func Up20170511195623(tx *sql.Tx) error {
	_, err := tx.Exec(`create table residents (
	id int not null auto_increment,
	firstname varchar(255),
	middlename varchar(255),
	lastname varchar(255),
	primary key (id)
)`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`create table units (
	id int not null auto_increment,
	name varchar(100),
	primary key (id)
)`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`create table units_residents (
	id int not null auto_increment,
	unit int not null,
	resident int not null,
	constraint fk_unit foreign key (unit)
		references units (id)
		on delete cascade,
	constraint fk_resident foreign key (resident)
		references residents (id)
		on delete cascade,
	primary key (id)
)`)
	if err != nil {
		return err
	}

	return nil
}

// Down20170511195623 downs
func Down20170511195623(tx *sql.Tx) error {
	_, err := tx.Exec(`drop table units_residents`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`drop table units`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`drop table residents`)
	if err != nil {
		return err
	}

	return nil
}
