package main

import (
	"log"

	"github.com/joho/godotenv"
)

const port = ":8080"

// init is invoked before main() to load .env
func init() {
	// loads values from .env into the system
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := NewConfig()

	// dataSourceName will be using in the data base connection
	dsn := conf.FormDSN()

	// create a connection to database
	conn, err := ConnectSQL(dsn)
	log.Println("Connecting to database...")

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	psql := NewPostgresRepo(conn)
	defer psql.DB.Close()

	httpServer := NewAPIserver(port, psql)
	httpServer.Run()
}
