package container

import (
	"database/sql"
	"log"

	_ "github.com/bmizerany/pq"
)

var dbS *sql.DB

const (
	DB_DRIVER = "postgres"
	DB_SOURCE = "host=localhost port=5432 user=root password=secret dbname=blog sslmode=disable"
)

func getDatabase() *sql.DB {
	var err error
	dbS, err = sql.Open(DB_DRIVER, DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	return dbS
}
