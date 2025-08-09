package seeders

import (
	"log"
	"split-bill-service/internal/models"

	"gorm.io/gorm"
)

func SeedUserFriends(db *gorm.DB) {
    log.Println("Seeding user friends...")

    userFriends := []models.UserFriend{
        {UserId: 1, FriendId: 2},
        {UserId: 2, FriendId: 1},
    }

    if err := db.Create(&userFriends).Error; err != nil {
        log.Fatalf("could not seed users: %v", err)
    }

}
