package extractor

import (
	"encoding/json"

	extractor "github.com/bhavyavarshney-123/catalyst/internal/extractor/models"
	"github.com/bhavyavarshney-123/catalyst/internal/models"
	"github.com/bhavyavarshney-123/catalyst/internal/services/ai"
)

type Extractor struct {
	ai ai.AIService
}

func NewExtractor(ai ai.AIService) *Extractor {
	return &Extractor{
		ai: ai,
	}
}

func (e *Extractor) ExtractOpportunity(email models.Email) (extractor.ExtractOpportunity, error) {

	content := buildContent(email)
	response, err := e.ai.Generate(OpportunityExtractionPrompt, content)

	if err != nil {
		return extractor.ExtractOpportunity{}, err
	}

	var opportunity extractor.ExtractOpportunity
	err = json.Unmarshal([]byte(response), &opportunity)
	if err != nil {
		return extractor.ExtractOpportunity{}, err
	}

	return opportunity, nil

}
