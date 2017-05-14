package main

import (
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/liszt-code/liszt/migrations"
)

var (
	driverName string
	dsn        string
)

func init() {
	flag.StringVar(&driverName, "driver", "mysql", "driver name")
	flag.StringVar(&dsn, "dsn", "root:@/test", "datasourcename uri")
}

func main() {
	flag.Parse()
	db, err := sqlx.Open(driverName, dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = migrations.Migrate(db.DB)
	if err != nil {
		log.Fatalln(err)
	}
}
