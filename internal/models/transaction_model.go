package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UniqueId string `gorm:"not null;unique"`
	Title string `gorm:"not null"`
	UserId uint `gorm:"not null"`
	StatusId uint `gorm:"not null"`
	ServiceCharge float64 `gorm:"not null default:0"`
	TaxCharge float64 `gorm:"not null default:0"`
	TotalPrice float64 `gorm:"not null default:0"`
	TransactionItems []TransactionItem `gorm:"foreignKey:TransactionId" json:"TransactionItems,omitempty"`
}