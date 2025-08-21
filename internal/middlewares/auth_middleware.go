package middlewares

import (
	"net/http"
	"split-bill-service/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization");	
		if authHeader == "" {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		tokenString := authHeaderParts[1]

		token, err := utils.VerifyJWTUserToken(tokenString)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

		c.Set("User", token.User)
		c.Next()
	}
}
