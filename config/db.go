package config

import (
	"github.com/go-pg/pg/v9"
	"log"
	"os"

	controllers "github.com/Witaly3/golang-rest/controllers"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     "db:5432",
		Database: "postgres",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)
	return db
}
