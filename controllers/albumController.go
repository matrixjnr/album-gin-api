package controllers

import (
	"john.com/album-gin-api/middleware"
	"john.com/album-gin-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAlbums(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var albums []models.Album
		db.Find(&albums)
		c.JSON(200, albums)
	}
}

func GetAlbumByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var album models.Album
		if result := db.First(&album, "id = ?", id); result.Error != nil {
			c.JSON(404, gin.H{"message": "album not found"})
			return
		}
		c.JSON(200, album)
	}
}

func PostAlbums(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authenticate the request
		if err := middleware.Authenticate(c); err != nil {
			c.JSON(401, gin.H{"message": "unauthorized"})
			return
		}

		var newAlbum models.Album
		if err := c.BindJSON(&newAlbum); err != nil {
			return
		}
		db.Create(&newAlbum)
		c.JSON(201, newAlbum)
	}
}

func DeleteAlbumByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authenticate the request
		if err := middleware.Authenticate(c); err != nil {
			c.JSON(401, gin.H{"message": "unauthorized"})
			return
		}

		id := c.Param("id")
		if result := db.Delete(&models.Album{}, "id = ?", id); result.Error != nil {
			c.JSON(404, gin.H{"message": "album not found"})
			return
		}
		c.JSON(200, gin.H{"message": "album deleted"})
	}
}
