package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is a standard format for API responses.

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func send(c *gin.Context, code int, data any, message string, err error) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	c.JSON(code, Response{
		Code:    code,
		Data:    data,
		Message: message,
		Error:   errMsg,
	})
}

func Success(c *gin.Context, data any, message string) {
	send(c, http.StatusOK, data, message, nil)
}

func Error(c *gin.Context, code int, err error) {
	send(c, code, nil, "", err)
}