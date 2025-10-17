package handlers

import (
	"net/http"
	"split-bill-service/internal/services"
	"split-bill-service/utils"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService *services.ProfileService
}

func NewProfileHandler(profileService *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: profileService}
}

func (h *ProfileHandler) ChangePassword(c *gin.Context) {
	var req struct {
		CurrentPassword string `json:"current_password" binding:"required"`
		NewPassword     string `json:"new_password" binding:"required"`
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

	err = h.profileService.ChangePassword(req.CurrentPassword, req.NewPassword, user.Email)
	if err != nil {
		c.Error(err)
		return
	}

	utils.SendResponse(c, http.StatusOK, nil, "Change password successful")
}
