package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"Code"`
	Data    any    `json:"Data,omitempty"`
	Message string `json:"Message"`
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
