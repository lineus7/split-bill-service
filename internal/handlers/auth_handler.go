package handlers

import (
	"net/http"
	"split-bill-service/internal/dtos"
	"split-bill-service/internal/services"
	"split-bill-service/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	user, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.Error(err)
		return
	}

	token, err := utils.GenerateJWTUserToken(user, time.Hour*24*30)
	if err != nil {
		c.Error(err)
		return
	}

	utils.SendResponse(c, http.StatusOK, dtos.ResponseLoginDTO{User: *user, AccessToken: token}, "Login successful")
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Username string `json:"username" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	_, err := h.authService.Register(req.Email, req.Password, req.Name, req.Username)
	if err != nil {
		c.Error(err)
		return
	}

	utils.SendResponse(c, http.StatusOK, nil, "Register successful")
}