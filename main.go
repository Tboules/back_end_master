package main

import (
	"context"

	"github.com/Tboules/back_end_master/api"
	db "github.com/Tboules/back_end_master/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const (
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Server was unable to start:", err)
	}
}
