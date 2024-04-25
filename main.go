package main

import (
	"database/sql"
	"log"

	"github.com/vinishs59/simplebank/api"
	db "github.com/vinishs59/simplebank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgressql://root:password@localhost:5432/simple_bank?sslmode=disable"
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