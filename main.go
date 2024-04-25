package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vinishs59/simplebank/api"
	db "github.com/vinishs59/simplebank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver,dbSource)
	log.Println(dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db", err )
	}else {
		log.Print("Successfully connected")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err !=nil {
		log.Fatal("cannot start the server",err)
	}else {
		log.Print("Server Started ....")
	}

}