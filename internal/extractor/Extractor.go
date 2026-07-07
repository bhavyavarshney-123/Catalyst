package extractor

import (
	extractor "github.com/bhavyavarshney-123/catalyst/internal/extractor/models"
	"github.com/bhavyavarshney-123/catalyst/internal/models"
	"github.com/bhavyavarshney-123/catalyst/internal/services/ai"
)

type Extractor struct {
	ai     ai.AIService
	prompt string
}

func NewExtractor(ai ai.AIService, prompt string) *Extractor {
	return &Extractor{
		ai:     ai,
		prompt: prompt,
	}
}

func (e *Extractor) ExtractOpportunity(email models.Email) (extractor.ExtractOpportunity, error) {
	return extractor.ExtractOpportunity{}, nil
}
