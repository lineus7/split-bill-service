package models

import (
	"gorm.io/gorm"
)

type TransactionItem struct {
	gorm.Model
	TransactionId uint `gorm:"not null"`
	ItemName string `gorm:"not null"`
	Price float64 `gorm:"not null default:0"`
	Transaction Transaction `gorm:"foreignKey:TransactionId"`
}