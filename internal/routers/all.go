package routers

import (
	"github.com/gin-gonic/gin"

	"split-bill-service/config"
	"split-bill-service/internal/middlewares"
	"split-bill-service/internal/repositories"
)

func SetupRouters(router *gin.Engine,connection *config.Connection, repos *repositories.RepositoriesSet) {
	SetupAuthRoutes(router,connection,repos)
	{
		r := router.Group("/api")
		r.Use(middlewares.AuthMiddleware())
		SetupReceiptRoutes(r,connection,repos)
	}
}
