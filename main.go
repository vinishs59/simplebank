package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vinishs59/simplebank/api"
	db "github.com/vinishs59/simplebank/db/sqlc"
	"github.com/vinishs59/simplebank/util"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	config, err :=util.Load(".")

	if err != nil {
		log.Fatal("Cannot load config")
	}

	log.Println("DB Driver:", config.DBDriver)

	conn, err := sql.Open(config.DBDriver,config.DBSource)
	log.Println(dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db", err )
	}else {
		log.Print("Successfully connected")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err !=nil {
		log.Fatal("cannot start the server",err)
	}else {
		log.Print("Server Started ....")
	}

}