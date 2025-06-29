package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"split-bill-service/internal/handlers"
	"split-bill-service/internal/repositories"
	"split-bill-service/internal/services"
)

func SetupAuthRoutes(router *gin.Engine, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	authHandler := handlers.NewAuthHandler(authService)

	api := router.Group("/auth")
	{
		api.POST("/login", authHandler.Login)
		api.POST("/register", authHandler.Register)
	}
}