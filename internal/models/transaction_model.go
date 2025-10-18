package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserId uint `gorm:"not null"`
	StatusId uint `gorm:"not null"`
	ServiceCharge float64 `gorm:"not null default:0"`
	TaxCharge float64 `gorm:"not null default:0"`
	TotalPrice float64 `gorm:"not null default:0"`
	User User `gorm:"foreignKey:UserId" json:"omitempty"`
	Status TransactionStatus `gorm:"foreignKey:StatusId" json:"omitempty"`
}