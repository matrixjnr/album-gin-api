package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"john.com/album-gin-api/models"
	"john.com/album-gin-api/utils"
)

func RegisterUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"message": "invalid request"})
			return
		}

		// Hash the password
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error"})
			return
		}
		user.Password = hashedPassword

		db.Create(&user)
		c.JSON(201, user)
	}
}

func LoginUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials models.User
		if err := c.BindJSON(&credentials); err != nil {
			c.JSON(400, gin.H{"message": "invalid request"})
			return
		}

		var user models.User
		if result := db.Where("email = ?", credentials.Email).First(&user); result.Error != nil {
			c.JSON(401, gin.H{"message": "unauthorized"})
			return
		}

		if err := utils.CheckPasswordHash(credentials.Password, user.Password); err != nil {
			c.JSON(401, gin.H{"message": "unauthorized"})
			return
		}

		token, err := utils.GenerateJWT(user.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		c.JSON(200, gin.H{"token": token})
	}
}

func GetUserProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(401, gin.H{"message": "unauthorized"})
			return
		}

		var user models.User
		if result := db.First(&user, userID); result.Error != nil {
			c.JSON(404, gin.H{"message": "user not found"})
			return
		}

		c.JSON(200, user)
	}
}
