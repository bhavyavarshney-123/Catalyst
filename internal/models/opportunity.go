package models

import (
	"time"

	"github.com/pgvector/pgvector-go"
)

type Opportunity struct {
	ID                int64      `json:"id"`
	CompanyName       string     `json:"company_name"`
	RoleApplied       string     `json:"role_applied"`
	ApplicationStatus string     `json:"application_status"`
	NextStepTodo      string     `json:"next_step_todo"`
	InterviewDate     *time.Time `json:"interview_date"`
	MeetingLink       string     `json:"meeting_link"`
	TestLink          string     `json:"test_link"`
	Comments          string     `json:"comments"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`

	Embedding pgvector.Vector `json:"embedding"`
}
