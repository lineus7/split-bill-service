package routers

import (
	"github.com/gin-gonic/gin"

	"split-bill-service/config"
	"split-bill-service/internal/middlewares"
	"split-bill-service/internal/repositories"
)

func SetupRouters(router *gin.Engine, connection *config.Connection, repos *repositories.RepositoriesSet) {
	// PUBLIC ROUTES
	SetupAuthRoutes(router, connection, repos)
	SetupPublicRoutes(router, connection, repos)
	// PROTECTED ROUTES
	{
		r := router.Group("")
		r.Use(middlewares.AuthMiddleware())
		SetupReceiptRoutes(r,connection,repos)
		SetupProfileRoutes(r,connection,repos)
		SetupTransactionRoutes(r,connection,repos)
	}
}
