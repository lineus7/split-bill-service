package config

import "gorm.io/gorm"

type Connection struct {
	DB *gorm.DB
}

func NewConnection(db *gorm.DB) *Connection {
	return &Connection{DB: db}
}
