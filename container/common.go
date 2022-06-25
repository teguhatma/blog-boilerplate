package container

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/teguhatma/blog-boilerplate/utils"
)

var (
	dbS *sql.DB
	err error
)

func getDatabase() *sql.DB {
	driver, source := utils.DatabaseVariable()
	dbS, err = sql.Open(driver, source)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	return dbS
}
