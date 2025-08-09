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

    seeders.SeedUsers(db)
    seeders.SeedTransactionStatus(db)
    seeders.SeedTransactions(db)

    log.Println("Seeding complete!")
}