package models

type UserFriend struct {
	ID uint `gorm:"primarykey"`
	UserId uint `gorm:"not null"`
	FriendId uint `gorm:"not null"`
	User User `gorm:"foreignKey:UserId"`
	Friend User `gorm:"foreignKey:FriendId"`
}