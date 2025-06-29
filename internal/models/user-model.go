package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"Name"`
	Email    string `json:"Email" gorm:"unique"`
	Password string `json:"Password" gorm:"not null"`
}