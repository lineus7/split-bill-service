package handlers

import (
	"net/http"
	"split-bill-service/internal/services"
	"split-bill-service/utils"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
}

func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) GetListTransaction (c *gin.Context) {
	user, err := utils.GetUserSession(c)
	if err != nil {
		c.Error(err)
		return
	}

	transactions, err := h.transactionService.GetListByUserId(user.ID)
	if err != nil {
		c.Error(err)
		return
	}

	utils.SendResponse(c, http.StatusOK, transactions, "Get list transaction successful")
}
