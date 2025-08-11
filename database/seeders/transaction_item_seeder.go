package seeders

import (
	"log"
	"split-bill-service/internal/models"

	"gorm.io/gorm"
)

func SeedTransactionItems(db *gorm.DB) {
    log.Println("Seeding transaction items...")
    
    transactionItems := []models.TransactionItem{
        {TransactionId: 1, ItemName: "Item 1", Price: 2000},
        {TransactionId: 1, ItemName: "Item 2", Price: 1000},
        {TransactionId: 1, ItemName: "Item 3", Price: 500},
    }

    if err := db.Create(&transactionItems).Error; err != nil {
        log.Fatalf("could not seed transaction items: %v", err)
    }
}
