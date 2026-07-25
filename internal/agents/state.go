package agents

import "github.com/bhavyavarshney-123/catalyst/internal/models"

type State struct {
	UserQuestion string

	Limit int

	Emails []models.Email

	Opportunities []models.Opportunity

	SelectedOpportunities []models.Opportunity

	DraftEmails []string

	Summary string

	Response string

	NeedsSync bool

	Err error
}
