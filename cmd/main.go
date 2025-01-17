package main

import (
	"log"

	"github.com/ihajar/ecom-api/cmd/api"
	"github.com/ihajar/ecom-api/db"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error in loading .env file!")
	}
	db.NewPostgreSQLStorage()
}

func main() {
	// DB connection
	psqlDB, err := db.DBConn.DB()

	if err != nil {
		log.Panic("Error in PostgreSQL connection!")
	}

	defer psqlDB.Close()

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
