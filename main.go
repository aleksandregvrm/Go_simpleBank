package main

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/banking/api"
	db "example.com/banking/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://bankingGo2:bankingGo2@localhost:5433/bankingGo2?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
)

func main() {
	if dbDriver == "" || dbSource == "" {
		log.Fatal("DB_DRIVER and POSTGRES_SERVICE_URL must be set as environment variables")
	}
	fmt.Println("Something")

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Couldn't connect to the database")
		return
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}
