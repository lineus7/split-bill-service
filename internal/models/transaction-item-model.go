package models

import "gorm.io/gorm"

type TransactionItem struct {
	gorm.Model
	TransactionId uint `gorm:"not null"`
	UserId uint `gorm:"not null"`
	AlternativeCustomerName string `gorm:"not null"`
	ItemName string `gorm:"not null"`
	TotalPrice float64 `gorm:"not null default:0"`
	AddOn []AddOn `gorm:"not null default:[]"`
	User User `gorm:"foreignKey:UserId"`
	Transaction Transaction `gorm:"foreignKey:TransactionId"`
}

type AddOn struct {
	Name string `gorm:"not null"`
	Price float64 `gorm:"not null default:0"`
}