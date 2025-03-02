package routes

import (
	"john.com/album-gin-api/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/albums", controllers.GetAlbums(db))
	router.GET("/albums/:id", controllers.GetAlbumByID(db))
	router.POST("/albums", controllers.PostAlbums(db))
	router.DELETE("/albums/:id", controllers.DeleteAlbumByID(db))
}
