package seeders

import (
	"log"
	"split-bill-service/internal/models"

	"gorm.io/gorm"
)

func SeedTransactionItemUsers(db *gorm.DB) {
    log.Println("Seeding transaction item users...")

	userId1 := uint(1)
	alternativeCustomerName1 := "Customer 1"
    alternativeCustomerName2 := "Customer 2"
    
    transactionItemUsers := []models.TransactionItemUser{
        {TransactionItemId: 1, UserId: &userId1},
        {TransactionItemId: 1, AlternativeCustomerName: &alternativeCustomerName1},
        {TransactionItemId: 2, AlternativeCustomerName: &alternativeCustomerName2},
    }

    if err := db.Create(&transactionItemUsers).Error; err != nil {
        log.Fatalf("could not seed transaction item users: %v", err)
    }
}
