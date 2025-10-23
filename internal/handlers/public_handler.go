package handlers

import (
	"net/http"
	"split-bill-service/internal/services"
	"split-bill-service/utils"

	"github.com/gin-gonic/gin"
)

type PublicHandler struct {
	publicService *services.PublicService
}

func NewPublicHandler(publicService *services.PublicService) *PublicHandler {
	return &PublicHandler{publicService: publicService}
}

func (h *PublicHandler) GetTransactionDetailByUniqueId(c *gin.Context) {
	uniqueId := c.Param("uniqueId")
	if uniqueId == "" {
		utils.SendResponse(c, http.StatusBadRequest, nil, "Unique id is required")
		return
	}

	transaction, err := h.publicService.GetTransactionDetailByUniqueId(uniqueId)
	if err != nil {
		c.Error(err)
		return
	}

	utils.SendResponse(c, http.StatusOK, transaction, "Get transaction detail successful")
}
