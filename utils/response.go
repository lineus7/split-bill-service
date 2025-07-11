package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
}

func SendResponse(c *gin.Context, code int, data any, message string) {
	c.JSON(code, Response{
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func SendErrorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, Response{
		Code:    code,
		Message: message,
	})
}
