package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var TestQueries *Queries
var TestDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	dbDriver := "postgres"
	dbSource := "postgres://bankingGo2:bankingGo2@localhost:5433/bankingGo2?sslmode=disable"

	if dbDriver == "" || dbSource == "" {
		log.Fatal("DB_DRIVER and POSTGRES_SERVICE_URL must be set as environment variables")
	}

	TestDb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Couldn't connect to the database")
		return
	}

	TestQueries = New(TestDb)
	os.Exit(m.Run())
}
