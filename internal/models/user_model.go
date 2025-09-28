package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"Name"`
	Username string `json:"Username" gorm:"unique"`
	Email    string `json:"Email" gorm:"unique"`
	Password string `json:"Password,omitempty" gorm:"not null"`
	RoleID   uint   `json:"RoleID" gorm:"not null"`
	Role     Role   `gorm:"foreignKey:RoleID;references:ID"`
}