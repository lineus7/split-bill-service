package models

type TransactionStatus struct {
	ID	uint `gorm:"primarykey"`
	Status string `gorm:"not null;unique"`
}