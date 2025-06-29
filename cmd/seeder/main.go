package main

import (
	"log"
	"split-bill-service/config"
	"split-bill-service/database"
	"split-bill-service/database/seeders"
)

func main() {
    cfg := config.LoadConfig()

    database.ConnectDB(cfg)
    db := database.DB

    log.Println("Seeding started...")

    seeders.SeedUsers(db)

    log.Println("Seeding complete!")
}