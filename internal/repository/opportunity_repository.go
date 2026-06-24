package repository

import (
	"database/sql"
	"fmt"

	"github.com/bhavyavarshney-123/catalyst/internal/models"
)

type OpportunityRepository struct {
	db *sql.DB
}

func NewOpportunityRepository(db *sql.DB) *OpportunityRepository {
	return &OpportunityRepository{
		db: db,
	}
}

func (r *OpportunityRepository) CreateOpportunity(opportunity models.Opportunity) error {

	query := `INSERT INTO opportunities (
    company_name,
    role_applied,
    application_status,
    next_step_todo,
    interview_date,
    meeting_link,
    test_link,
    comments
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)`

	_, err := r.db.Exec(query, opportunity.CompanyName,
		opportunity.RoleApplied,
		opportunity.ApplicationStatus,
		opportunity.NextStepTodo,
		opportunity.InterviewDate,
		opportunity.MeetingLink,
		opportunity.TestLink,
		opportunity.Comments)

	if err != nil {
		return err
	}

	return nil
}

func (r *OpportunityRepository) GetOpportunities() ([]models.Opportunity, error) {

	query := `SELECT * FROM opportunities`
	var opportunities []models.Opportunity

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var opportunity models.Opportunity

		err := rows.Scan(
			&opportunity.ID,
			&opportunity.CompanyName,
			&opportunity.RoleApplied,
			&opportunity.ApplicationStatus,
			&opportunity.NextStepTodo,
			&opportunity.InterviewDate,
			&opportunity.MeetingLink,
			&opportunity.TestLink,
			&opportunity.Comments,
			&opportunity.CreatedAt,
			&opportunity.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		opportunities = append(opportunities, opportunity)
	}

	return opportunities, nil
}

func (r *OpportunityRepository) GetOpportunityByID(id int64) (*models.Opportunity, error) {

	query := `SELECT * FROM opportunities WHERE id=$1`
	var opportunity models.Opportunity

	row := r.db.QueryRow(query, id)

	err := row.Scan(
		&opportunity.ID,
		&opportunity.CompanyName,
		&opportunity.RoleApplied,
		&opportunity.ApplicationStatus,
		&opportunity.NextStepTodo,
		&opportunity.InterviewDate,
		&opportunity.MeetingLink,
		&opportunity.TestLink,
		&opportunity.Comments,
		&opportunity.CreatedAt,
		&opportunity.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &opportunity, nil
}

func (r *OpportunityRepository) UpdateOpportunity(opportunity models.Opportunity, id int64) error {

	query := `UPDATE opportunities
SET
    application_status = $1,
    next_step_todo = $2,
    interview_date = $3,
    meeting_link = $4,
    test_link = $5,
    comments = $6
WHERE id = $7`

	results, err := r.db.Exec(query,
		opportunity.ApplicationStatus,
		opportunity.NextStepTodo,
		opportunity.InterviewDate,
		opportunity.MeetingLink,
		opportunity.TestLink,
		opportunity.Comments,
		id,
	)
	if err != nil {
		return err
	}

	rowAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return fmt.Errorf("opportunity not found to be updated")
	}

	return nil
}

func (r *OpportunityRepository) DeleteOpportunity(id int64) error {

	query := `DELETE FROM opportunities WHERE id = $1`
	results, err := r.db.Exec(query,
		id,
	)
	if err != nil {
		return err
	}

	rowAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return fmt.Errorf("opportunity not found to be deleted")
	}

	return nil
}
