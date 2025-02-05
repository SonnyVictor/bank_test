package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnyvictok/bank_test/api"
	db "github.com/sonnyvictok/bank_test/db/sqlc"
	"github.com/sonnyvictok/bank_test/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Error connect DB", err)
	}
	storeDB := db.NewStore(conn)
	server, err := api.NewServer(config, storeDB)
	if err != nil {
		log.Fatal("Server is not connect ", err)
	}

	fmt.Println("Address Server", config.ServerAddress)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
