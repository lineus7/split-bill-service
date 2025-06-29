package seeders

import (
	"log"
	"split-bill-service/internal/models"
	"split-bill-service/utils"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
    log.Println("Seeding users...")

    password, _ := utils.HashPassword("password")

    users := []models.User{
        {Name: "Alice", Email: "alice@example.com", Password: password},
        {Name: "Bob", Email: "bob@example.com", Password: password},
    }

    if err := db.Create(&users).Error; err != nil {
        log.Fatalf("could not seed users: %v", err)
    }

}
