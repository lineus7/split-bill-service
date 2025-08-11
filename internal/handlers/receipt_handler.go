package handlers

import (
	"net/http"
	"split-bill-service/internal/services"
	"split-bill-service/utils"

	"github.com/gin-gonic/gin"
)

type ReceiptHandler struct {
	receiptService *services.ReceiptService
}

func NewReceiptHandler(receiptService *services.ReceiptService) *ReceiptHandler {
	return &ReceiptHandler{receiptService: receiptService}
}

func (h *ReceiptHandler) GenerateReceiptJSON(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.Error(err)
		return
	}

	content, err := h.receiptService.GenerateReceiptJSON(fileHeader)
	if err != nil {
		c.Error(err)
		return
	}

	utils.SendResponse(c, http.StatusOK, content, "Generate receipt JSON successful")
}
	