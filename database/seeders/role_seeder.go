package seeders

import (
	"log"
	"split-bill-service/internal/models"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
    log.Println("Seeding roles...")

    roles := []models.Role{
        {RoleName: "User"},
        {RoleName: "SuperAdmin"},
    }

    if err := db.Create(&roles).Error; err != nil {
        log.Fatalf("could not seed roles: %v", err)
    }

}
