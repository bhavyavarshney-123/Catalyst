package extractor

import (
	extractor "github.com/bhavyavarshney-123/catalyst/internal/extractor/models"
	"github.com/bhavyavarshney-123/catalyst/internal/models"
	"github.com/bhavyavarshney-123/catalyst/internal/services/AI"
)

type Extractor struct {
	ai     AI.AIService
	prompt string
}

func NewExtractor(ai AI.AIService, prompt string) *Extractor {
	return &Extractor{
		ai:     ai,
		prompt: prompt,
	}
}

func (e *Extractor) ExtractOpportunity(email models.Email) (extractor.ExtractOpportunity, error) {

}
