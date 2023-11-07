package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lysenkopavlo/goperson/internal/config"
	"github.com/lysenkopavlo/goperson/internal/repo/driver"
	"github.com/lysenkopavlo/goperson/internal/repo/storage"
)

const port = ":8080"

// init is invoked before main() to load .env
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.NewConfig()

	// dataSourceName will be using in the data base connection
	dsn := conf.FormDSN()

	// create a connection to database
	conn, err := driver.ConnectSQL(dsn)
	log.Println("Connecting to database...")

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	psql := storage.NewPostgresRepo(conn)
	defer psql.DB.Close()

	httpServer := NewAPIserver(port, psql, conf.EnrichSource)
	httpServer.Run()

}
