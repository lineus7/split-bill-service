package seeders

import (
	"log"
	"split-bill-service/internal/models"

	"gorm.io/gorm"
)

func SeedTransactionItemAddOns(db *gorm.DB) {
    log.Println("Seeding transaction item add-ons...")

    transactionItemAddOns := []models.TransactionItemAddOn{
        {TransactionItemId: 1, ItemName: "Add On 1", Price: 2000},
        {TransactionItemId: 1, ItemName: "Add On 2", Price: 1000},
        {TransactionItemId: 2, ItemName: "Add On 3", Price: 500},
    }

    if err := db.Create(&transactionItemAddOns).Error; err != nil {
        log.Fatalf("could not seed transaction item add-ons: %v", err)
    }
}
