package main

import (
	"database/sql"
	"log"
	"os"

	"example.com/banking/api"
	db "example.com/banking/db/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DBDriver := os.Getenv("DB_DRIVER")
	DBSource := os.Getenv("DB_SOURCE")
	ServerAddress := os.Getenv("SERVER_ADDRESS")
	conn, err := sql.Open(DBDriver, DBSource)
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}
