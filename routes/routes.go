package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"john.com/album-gin-api/controllers"
	"john.com/album-gin-api/middleware"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/albums", controllers.GetAlbums(db))
	router.GET("/albums/:id", controllers.GetAlbumByID(db))
	router.POST("/albums", controllers.PostAlbums(db))
	router.DELETE("/albums/:id", controllers.DeleteAlbumByID(db))

	router.POST("/register", controllers.RegisterUser(db))
	router.POST("/login", controllers.LoginUser(db))

	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", controllers.GetUserProfile(db))
	}
}
