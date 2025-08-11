package services

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"split-bill-service/internal/repositories"
	"split-bill-service/prompts"
	"split-bill-service/utils"
	"strings"
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
	
	resultString, err := s.geminiRepository.GenerateContentWithBytes(fileBytes, prompts.GenerateReceiptPrompt)
	if err != nil {
		return "", err
	}

	resultString = strings.TrimPrefix(resultString, "```json\n")
	resultString = strings.TrimSuffix(resultString, "\n```")

	var jsonMap map[string]interface{}
    if err := json.Unmarshal([]byte(resultString), &jsonMap); err != nil {
        return "", fmt.Errorf("failed to parse JSON: %w", err)
    }

    if _, ok := jsonMap["error"]; ok {
        return "", utils.NewAppError(http.StatusBadRequest, jsonMap["error"].(string), nil)
    }

	return resultString, nil
}
