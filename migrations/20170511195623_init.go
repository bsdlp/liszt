package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up, down)
}

func up(tx *sql.Tx) error {
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
	unit int not null references units (id),
	resident int not null references residents (id),
	primary key (unit, resident)
)`)
	if err != nil {
		return err
	}

	return nil
}

func down(tx *sql.Tx) error {
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
