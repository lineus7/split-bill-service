package models

import (
	"gorm.io/gorm"
)

type TransactionItem struct {
	gorm.Model
	TransactionId uint `gorm:"not null"`
	UserId *uint 
	AlternativeCustomerName *string 
	ItemName string `gorm:"not null"`
	Price float64 `gorm:"not null default:0"`
	User User `gorm:"foreignKey:UserId"`
	Transaction Transaction `gorm:"foreignKey:TransactionId"`
}