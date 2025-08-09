package middlewares

import (
	"errors"
	"net/http"
	"split-bill-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var appErr *utils.AppError

			var validationErr validator.ValidationErrors
			if errors.As(err, &validationErr) {
				utils.SendErrorResponse(c, http.StatusBadRequest, validationErr.Error())
				return
			}

			if errors.As(err, &appErr) {
				utils.SendErrorResponse(c, appErr.Code, appErr.Message)
				return
			}
			
			utils.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
	}
}
