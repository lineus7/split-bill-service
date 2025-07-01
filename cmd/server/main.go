package main

import (
	"fmt"
	"log"
	"split-bill-service/config"
	"split-bill-service/database"
	"split-bill-service/internal/middlewares"
	"split-bill-service/internal/routers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	cfg := config.LoadConfig()
	database.ConnectDB(cfg)
	db = database.DB
	
	gin.SetMode(cfg.GinMode)
	r := gin.Default()

	r.Use(middlewares.ErrorHandlingMiddleware())
	routers.SetupAuthRoutes(r, db)

	log.Printf("Server starting on port %s", cfg.HTTPPort)
	if err := r.Run(fmt.Sprintf(":%s", cfg.HTTPPort)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}