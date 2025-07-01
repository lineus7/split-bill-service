package middlewares

import (
	"errors"
	"net/http"
	"split-bill-service/utils"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var appErr *utils.AppError
			if errors.As(err, &appErr) {
				utils.SendResponse(c, appErr.Code, nil, appErr.Message)
			} else {
				utils.SendResponse(c, http.StatusInternalServerError, nil, "Internal Server Error")
			}
		}
	}
}
