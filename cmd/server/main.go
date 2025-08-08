package main

import (
	"fmt"
	"log"
	"split-bill-service/config"
	"split-bill-service/database"
	"split-bill-service/internal/middlewares"
	"split-bill-service/internal/repositories"
	"split-bill-service/internal/routers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	cfg := config.LoadConfig()
	db = database.ConnectDB(cfg)
	connection := config.NewConnection(db)
	repos := repositories.NewRepositoriesSet(connection)
	
	gin.SetMode(cfg.GinMode)
	r := gin.Default()

	r.Use(middlewares.ErrorHandlingMiddleware())
	routers.SetupRouters(r, connection, repos)

	log.Printf("Server starting on port %s", cfg.HTTPPort)
	if err := r.Run(fmt.Sprintf(":%s", cfg.HTTPPort)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}