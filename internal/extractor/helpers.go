package extractor

import (
	"fmt"
	"time"

	extractor "github.com/bhavyavarshney-123/catalyst/internal/extractor/models"
	"github.com/bhavyavarshney-123/catalyst/internal/models"
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
	return models.Opportunity{
		CompanyName:       extracted.CompanyName,
		RoleApplied:       extracted.RoleApplied,
		ApplicationStatus: extracted.ApplicationStatus,
		NextStepTodo:      extracted.NextStepTodo,
		InterviewDate:     &interviewTime,
		MeetingLink:       extracted.MeetingLink,
		TestLink:          extracted.TestLink,
		Comments:          extracted.Comments,
		Embedding:         embedding,
	}, nil
}
