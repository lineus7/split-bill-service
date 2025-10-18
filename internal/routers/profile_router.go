package routers

import (
	"github.com/gin-gonic/gin"

	"split-bill-service/config"
	"split-bill-service/internal/handlers"
	"split-bill-service/internal/repositories"
	"split-bill-service/internal/services"
)

func SetupProfileRoutes(router *gin.RouterGroup,connection *config.Connection, repos *repositories.RepositoriesSet) {
	profileService := services.NewProfileService(repos)
	profileHandler := handlers.NewProfileHandler(profileService)

	api := router.Group("/profile")
	{
		api.POST("/change-password", profileHandler.ChangePassword)
	}
}
