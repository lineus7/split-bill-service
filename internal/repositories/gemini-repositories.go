package repositories

import (
	"context"
	"log"
	"split-bill-service/config"

	"google.golang.org/genai"
)

func NewGeminiRepository(connection *config.Connection) *GeminiRepository {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &GeminiRepository{
		client: client,
	}
}

type GeminiRepository struct {
	client *genai.Client
}

func (r *GeminiRepository) GenerateContent(model string, prompt string) (string, error) {
	result, err := r.client.Models.GenerateContent(
		context.Background(),
		model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", err
	}
	return result.Text(), nil
}