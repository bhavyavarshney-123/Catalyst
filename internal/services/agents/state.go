package services

import "github.com/bhavyavarshney-123/catalyst/internal/models"

type State struct {
	UserQuestion string

	Emails []models.email

	Opportunities []models.Opportunity

	SelectedOpportunities []models.Opportunity

	DraftEmails []string

	Summary string

	Response string

	Err error
}
