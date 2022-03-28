package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/teguhatma/blog-boilerplate/repository"
)

var testQueries *repository.Queries
var testDB *sql.DB

const (
	DB_DRIVER = "postgres"
	DB_SOURCE = "postgresql://root:secret@localhost:5432/blog?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(DB_DRIVER, DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = repository.New(testDB)

	os.Exit(m.Run())
}
