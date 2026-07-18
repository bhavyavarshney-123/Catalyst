package extractor

import (
	"fmt"
	"time"

	extractor "github.com/bhavyavarshney-123/catalyst/internal/extractor/models"
	"github.com/bhavyavarshney-123/catalyst/internal/models"
	"github.com/pgvector/pgvector-go"
)

func buildContent(email models.Email) string {
	return fmt.Sprintf(
		`From: %s
         Subject: %s
		 Body:%s`,
		email.From,
		email.Subject,
		email.Body,
	)
}

func ToOpportunity(extracted extractor.ExtractOpportunity, embedding []float64) (models.Opportunity, error) {
	var interviewTime time.Time
	if extracted.InterviewDate != "" {
		t, err := time.Parse(time.RFC3339, extracted.InterviewDate)
		if err != nil {
			return models.Opportunity{}, err
		}
		interviewTime = t
	}

	vector := pgvector.NewVector(ToFloat32(embedding))
	return models.Opportunity{
		CompanyName:       extracted.CompanyName,
		RoleApplied:       extracted.RoleApplied,
		ApplicationStatus: extracted.ApplicationStatus,
		NextStepTodo:      extracted.NextStepTodo,
		InterviewDate:     &interviewTime,
		MeetingLink:       extracted.MeetingLink,
		TestLink:          extracted.TestLink,
		Comments:          extracted.Comments,
		Embedding:         vector,
		ExternalID:        extracted.ExternalID,
		Source:            extracted.Source,
	}, nil
}

func ToFloat32(values []float64) []float32 {
	result := make([]float32, len(values))

	for i, v := range values {
		result[i] = float32(v)
	}

	return result
}
