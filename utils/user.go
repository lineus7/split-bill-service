package utils

import (
	"errors"
	"split-bill-service/internal/models"

	"github.com/gin-gonic/gin"
)

func GetUserSession(c *gin.Context) (*models.User, error) {
	user, exists := c.Get("User")
	if !exists {
		return nil, errors.New("user not found")
	}
	return user.(*models.User), nil
}