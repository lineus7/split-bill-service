package dtos

import "split-bill-service/internal/models"

type ResponseLoginDTO struct {
	User models.User `json:"User"`
	AccessToken string `json:"Token"`
}