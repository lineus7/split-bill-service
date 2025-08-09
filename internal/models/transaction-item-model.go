package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TransactionItem struct {
	gorm.Model
	TransactionId uint `gorm:"not null"`
	UserId *uint 
	AlternativeCustomerName string 
	ItemName string `gorm:"not null"`
	TotalPrice float64 `gorm:"not null default:0"`
	AddOn datatypes.JSON `gorm:"default:'[]'"`
	User User `gorm:"foreignKey:UserId"`
	Transaction Transaction `gorm:"foreignKey:TransactionId"`
}