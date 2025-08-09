package routers

import (
	"github.com/gin-gonic/gin"

	"split-bill-service/config"
	"split-bill-service/internal/handlers"
	"split-bill-service/internal/repositories"
	"split-bill-service/internal/services"
)

func SetupAuthRoutes(router *gin.Engine,connection *config.Connection, repos *repositories.RepositoriesSet) {
	authService := services.NewAuthService(repos)
	authHandler := handlers.NewAuthHandler(authService)

	api := router.Group("/auth")
	{
		api.POST("/login", authHandler.Login)
		api.POST("/register", authHandler.Register)
		api.POST("/example", authHandler.Example)
	}
}