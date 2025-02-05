package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnyvictok/bank_test/util"
)

const (
	dbDriver = "postgresql"
	dbSource = "postgresql://root:secret@localhost:5000/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Error Load config ", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Error connect to DB", err)
	}
	defer connPool.Close()
	testQueries = New(connPool)
	testDB = connPool
	os.Exit(m.Run())

}
