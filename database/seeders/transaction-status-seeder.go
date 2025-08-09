package seeders

import (
	"log"
	"split-bill-service/internal/models"

	"gorm.io/gorm"
)

func SeedTransactionStatus(db *gorm.DB) {
    log.Println("Seeding transaction status...")

    transactionStatus := []models.TransactionStatus{
        {Status: "Pending"},
        {Status: "Completed"},
    }

    if err := db.Create(&transactionStatus).Error; err != nil {
        log.Fatalf("could not seed transaction status: %v", err)
    }
}
