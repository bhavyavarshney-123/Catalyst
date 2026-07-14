package rag

import (
	"strings"

	"github.com/bhavyavarshney-123/catalyst/internal/models"
)

func BuildContext(opportunities []models.Opportunity) string {
	var builder strings.Builder

	for _, opportunity := range opportunities {
		builder.WriteString("Company: ")
		builder.WriteString(opportunity.CompanyName)
		builder.WriteString("\n")

		builder.WriteString("Role: ")
		builder.WriteString(opportunity.RoleApplied)
		builder.WriteString("\n")

		builder.WriteString("Application Status: ")
		builder.WriteString(opportunity.ApplicationStatus)
		builder.WriteString("\n")

		builder.WriteString("Next Step: ")
		builder.WriteString(opportunity.NextStepTodo)
		builder.WriteString("\n")

		builder.WriteString("Interview Date: ")
		builder.WriteString(opportunity.InterviewDate.String())
		builder.WriteString("\n")

		builder.WriteString("Meeting Link: ")
		builder.WriteString(opportunity.MeetingLink)
		builder.WriteString("\n")

		builder.WriteString("Test Link: ")
		builder.WriteString(opportunity.TestLink)
		builder.WriteString("\n")

		builder.WriteString("Comments: ")
		builder.WriteString(opportunity.Comments)
	}
	return builder.String()
}
