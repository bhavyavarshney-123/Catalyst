package extractor

import (
	"fmt"

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

func ToOpportunity(extracted extractor.ExtractOpportunity) models.Opportunity {
	return models.Opportunity{
		CompanyName:       extracted.CompanyName,
		RoleApplied:       extracted.RoleApplied,
		ApplicationStatus: extracted.ApplicationStatus,
		NextStepTodo:      extracted.NextStepTodo,
		InterviewDate:     extracted.InterviewDate,
		MeetingLink:       extracted.MeetingLink,
		TestLink:          extracted.TestLink,
		Comments:          extracted.Comments,
	}
}
