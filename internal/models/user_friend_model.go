package models

import "gorm.io/gorm"

type UserFriend struct {
	gorm.Model
	UserId uint `gorm:"not null"`
	FriendId uint `gorm:"not null"`
	User User `gorm:"foreignKey:UserId"`
	Friend User `gorm:"foreignKey:FriendId"`
}