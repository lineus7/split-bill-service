package middlewares

import (
	"errors"
	"log"
	"net/http"
	"split-bill-service/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("AuthMiddleware")
		tokenString, err := c.Cookie("token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
				return
			}
			c.Error(err)
			return
		}

		token, err := utils.VerifyJWTUserToken(tokenString)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

		c.Set("User", token.User)
		c.Next()
	}
}
