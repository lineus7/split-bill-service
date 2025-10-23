package models

import (
	"gorm.io/gorm"
)

type TransactionItemAddOn struct {
	gorm.Model
	TransactionItemId uint `gorm:"not null"`
	ItemName string `gorm:"not null"`
	Price float64 `gorm:"not null default:0"`
}