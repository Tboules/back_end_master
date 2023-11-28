package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Tboules/back_end_master/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Could not load config:", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("TestMain cannot connect to db:", err)
	}

	testStore = NewStore(conn)

	os.Exit(m.Run())
}
