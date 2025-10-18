package routers

import (
	"github.com/gin-gonic/gin"

	"split-bill-service/config"
	"split-bill-service/internal/handlers"
	"split-bill-service/internal/repositories"
	"split-bill-service/internal/services"
)

func SetupTransactionRoutes(router *gin.RouterGroup,connection *config.Connection, repos *repositories.RepositoriesSet) {
	transactionService := services.NewTransactionService(repos)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	api := router.Group("/transaction")
	{
		api.GET("/list", transactionHandler.GetListTransaction)
	}
}
