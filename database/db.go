package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres dbname=go-db sslmode=disable host=localhost port=5432, Timezone=Asia/Jakarta",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connect to database: %v", err)
	}

	fmt.Println("Database Connect Successfully")
	return db
}
