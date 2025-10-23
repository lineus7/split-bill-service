package models

import (
	"gorm.io/gorm"
)

type TransactionItemUser struct {
	gorm.Model
	TransactionItemId uint `gorm:"not null"`
	UserId *uint
	AlternativeCustomerName *string 
	User User `gorm:"foreignKey:UserId" json:"User,omitempty"`
}