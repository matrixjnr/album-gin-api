package main

import (
	"john.com/album-gin-api/models"
	"john.com/album-gin-api/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	// Initialize the database
	db, err = gorm.Open(sqlite.Open("albums.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Album{})

	router := gin.Default()
	routes.RegisterRoutes(router, db)

	router.Run("localhost:8080")
}
