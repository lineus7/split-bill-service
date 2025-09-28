package models

type Role struct {
	ID        int    `json:"id" gorm:"primarykey"`
	RoleName  string `json:"role_name" gorm:"unique;not null"`
}