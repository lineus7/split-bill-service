package seeders

import (
	"log"
	"split-bill-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedTransactions(db *gorm.DB) {
    log.Println("Seeding transactions...")

    users := []models.Transaction{
        {UniqueId: uuid.New().String(), UserId: 1, StatusId: 1, Title: "Transaction 1", ServiceCharge: 1000, TaxCharge: 2000, TotalPrice: 100000},
    }

    if err := db.Create(&users).Error; err != nil {
        log.Fatalf("could not seed transactions: %v", err)
    }

}
