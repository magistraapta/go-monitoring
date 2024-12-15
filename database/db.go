package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func GetDb() (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=go-db sslmode=disable host=localhost port=5432")
	if err != nil {
		return nil, err
	}

	DB = db
	log.Println("Database connection established")
	return db, nil
}
