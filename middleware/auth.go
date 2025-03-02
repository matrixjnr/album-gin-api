package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"john.com/album-gin-api/utils"
)

func Authenticate(c *gin.Context) error {
	token := c.GetHeader("Authorization")
	if token == "" {
		return errors.New("no token provided")
	}

	_, err := utils.ValidateJWT(token)
	if err != nil {
		return err
	}

	return nil
}
