package main

import (
	"go-backend/database"

	user "go-backend/internal/model"

	"go-backend/internal/router"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	// Initialize db connection
	db = database.GetDb()

	db.AutoMigrate(&user.User{})

	r := router.SetupRouter(db)

	r.Run()
}
