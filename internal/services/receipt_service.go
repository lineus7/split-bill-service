package services

import (
	"io"
	"mime/multipart"
	"split-bill-service/internal/repositories"
	"split-bill-service/prompts"
)
func NewReceiptService(repos *repositories.RepositoriesSet) *ReceiptService {
	return &ReceiptService{
		geminiRepository: repos.GeminiRepository,
	}
}
type ReceiptService struct {
	geminiRepository *repositories.GeminiRepository
}

func (s *ReceiptService) GenerateReceiptJSON(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	
	return s.geminiRepository.GenerateContentWithBytes(fileBytes, prompts.GenerateReceiptPrompt)
}
