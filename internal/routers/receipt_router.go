package routers

import (
	"github.com/gin-gonic/gin"

	"split-bill-service/config"
	"split-bill-service/internal/handlers"
	"split-bill-service/internal/repositories"
	"split-bill-service/internal/services"
)

func SetupReceiptRoutes(router *gin.RouterGroup,connection *config.Connection, repos *repositories.RepositoriesSet) {
	receiptService := services.NewReceiptService(repos)
	receiptHandler := handlers.NewReceiptHandler(receiptService)

	api := router.Group("/receipt")
	{
		api.POST("/generate", receiptHandler.GenerateReceiptJSON)
	}
}