package main

import (
	"database/sql"
	"log"

	"github.com/dipu626/simple-bank/api"
	db "github.com/dipu626/simple-bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:6565@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err = server.Start(serverAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
