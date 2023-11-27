package main

import (
	"context"

	"github.com/Tboules/back_end_master/api"
	db "github.com/Tboules/back_end_master/db/sqlc"
	"github.com/Tboules/back_end_master/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Config files unable to load", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Server was unable to start:", err)
	}
}
