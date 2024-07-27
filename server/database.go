package server

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB *sqlx.DB
}

func SetupDatabase() Database {
	// config.ConnectionString
	db, err := sqlx.Connect("postgres", "user=postgres password=root dbname=doalivros sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return Database{
		DB: db,
	}
}
