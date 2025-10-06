package handlers

import (
	"net/http"
	"split-bill-service/internal/services"
	"split-bill-service/utils"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	authService *services.ProfileService
}

func NewProfileHandler(authService *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{authService: authService}
}

func (h *ProfileHandler) ChangePassword(c *gin.Context) {
	var req struct {
		OldPassword    string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	user, err := utils.GetUserSession(c)
	if err != nil {
		c.Error(err)
		return
	}

	user, err = h.authService.ChangePassword(user.Email, req.OldPassword, req.NewPassword)
	if err != nil {
		c.Error(err)
		return
	}

	utils.SendResponse(c, http.StatusOK, nil, "Change password successful")
}