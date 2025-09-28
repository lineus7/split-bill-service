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
        {Name: "Super Admin", Username: "superadmin", Email: "superadmin@example.com", Password: password, RoleID: 1, IsActive: true},
        {Name: "User", Username: "user", Email: "user@example.com", Password: password, RoleID: 2, IsActive: true},
    }

    if err := db.Create(&users).Error; err != nil {
        log.Fatalf("could not seed users: %v", err)
    }

}
