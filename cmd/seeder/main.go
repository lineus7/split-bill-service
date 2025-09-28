package main

import (
	"log"
	"split-bill-service/config"
	"split-bill-service/database"
	"split-bill-service/database/seeders"
)

func main() {
    cfg := config.LoadConfig()

    db := database.ConnectDB(cfg)

    log.Println("Seeding started...")

    seeders.SeedRoles(db)
    seeders.SeedUsers(db)
    seeders.SeedTransactionStatus(db)
    seeders.SeedTransactions(db)
    seeders.SeedTransactionItems(db)
    seeders.SeedTransactionItemAddOns(db)
    seeders.SeedUserFriends(db)
    seeders.SeedTransactionItemUsers(db)

    log.Println("Seeding complete!")
}