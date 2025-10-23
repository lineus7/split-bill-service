package routers

import (
	"github.com/gin-gonic/gin"

	"split-bill-service/config"
	"split-bill-service/internal/handlers"
	"split-bill-service/internal/repositories"
	"split-bill-service/internal/services"
)

func SetupPublicRoutes(router *gin.Engine,connection *config.Connection, repos *repositories.RepositoriesSet) {
	publicService := services.NewPublicService(repos)
	publicHandler := handlers.NewPublicHandler(publicService)

	api := router.Group("/public")
	{
		api.GET("/transaction-detail/:uniqueId", publicHandler.GetTransactionDetailByUniqueId)
	}
}
