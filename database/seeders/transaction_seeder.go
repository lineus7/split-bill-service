package seeders

import (
	"log"
	"split-bill-service/internal/models"

	"gorm.io/gorm"
)

func SeedTransactions(db *gorm.DB) {
    log.Println("Seeding transactions...")

    users := []models.Transaction{
        {UserId: 1, StatusId: 1, ServiceCharge: 1000, TaxCharge: 2000, TotalPrice: 100000},
        {UserId: 2, StatusId: 2, ServiceCharge: 2000, TaxCharge: 4000, TotalPrice: 100000},
    }

    if err := db.Create(&users).Error; err != nil {
        log.Fatalf("could not seed transactions: %v", err)
    }

}
