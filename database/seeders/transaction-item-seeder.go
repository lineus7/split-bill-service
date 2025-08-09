package seeders

import (
	"log"
	"split-bill-service/internal/models"

	"gorm.io/gorm"
)

func SeedTransactionItems(db *gorm.DB) {
    log.Println("Seeding transaction items...")

    transactionItems := []models.TransactionItem{
        {UserId: nil, TransactionId: 1, AlternativeCustomerName: "Customer 1", ItemName: "Item 1", TotalPrice: 2000},
        {UserId: nil, TransactionId: 2, AlternativeCustomerName: "Customer 2", ItemName: "Item 2", TotalPrice: 1000},
        {UserId: nil, TransactionId: 2, AlternativeCustomerName: "Customer 3", ItemName: "Item 3", TotalPrice: 500},
    }

    if err := db.Create(&transactionItems).Error; err != nil {
        log.Fatalf("could not seed transaction items: %v", err)
    }
}
