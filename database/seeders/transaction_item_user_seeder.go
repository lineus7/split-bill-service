package seeders

import (
	"log"
	"split-bill-service/internal/models"

	"gorm.io/gorm"
)

func SeedTransactionItemUsers(db *gorm.DB) {
    log.Println("Seeding transaction item users...")

	userId1 := uint(1)
	userId2 := uint(2)
	alternativeCustomerName1 := "Customer 1"
	
    transactionItemUsers := []models.TransactionItemUser{
        {TransactionItemId: 1, UserId: &userId1},
        {TransactionItemId: 1, UserId: &userId2},
        {TransactionItemId: 2, AlternativeCustomerName: &alternativeCustomerName1},
    }

    if err := db.Create(&transactionItemUsers).Error; err != nil {
        log.Fatalf("could not seed transaction item users: %v", err)
    }
}
