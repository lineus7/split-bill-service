package repositories

import (
	"context"
	"log"
	"split-bill-service/config"

	"google.golang.org/genai"
)

var model = "gemini-2.5-pro"

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

func (r *GeminiRepository) GenerateContent(prompt string) (string, error) {
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

func (r *GeminiRepository) GenerateContentWithBytes(imageBytes []byte, prompt string) (string, error) {
	parts := []*genai.Part{
		genai.NewPartFromBytes(imageBytes, "image/jpeg"),
		genai.NewPartFromText(prompt),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}
	
	result, err := r.client.Models.GenerateContent(
		context.Background(),
		model,
		contents,
		nil,
	)
	if err != nil {
		return "", err
	}

	return result.Text(), nil
}
	