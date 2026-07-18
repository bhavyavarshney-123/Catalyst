package extractor

type ExtractOpportunity struct {
	IsOpportunity     bool   `json:"is_opportunity"`
	CompanyName       string `json:"company_name"`
	RoleApplied       string `json:"role_applied"`
	ApplicationStatus string `json:"application_status"`
	NextStepTodo      string `json:"next_step_todo"`
	InterviewDate     string `json:"interview_date"`
	MeetingLink       string `json:"meeting_link"`
	TestLink          string `json:"test_link"`
	Comments          string `json:"comments"`

	Source     string `json:"source"`
	ExternalID string `json:"externalid"`
}
