package repository

import (
	"database/sql"
	"fmt"

	"github.com/bhavyavarshney-123/catalyst/internal/extractor"
	"github.com/bhavyavarshney-123/catalyst/internal/models"
	"github.com/pgvector/pgvector-go"
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
    comments,
	embedding,
	source,
	external_id

)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
	$9,
	$10,
	$11
)
ON CONFLICT (source, external_id)
DO NOTHING`

	_, err := r.db.Exec(query, opportunity.CompanyName,
		opportunity.RoleApplied,
		opportunity.ApplicationStatus,
		opportunity.NextStepTodo,
		opportunity.InterviewDate,
		opportunity.MeetingLink,
		opportunity.TestLink,
		opportunity.Comments,
		opportunity.Embedding,
		opportunity.Source,
		opportunity.ExternalID)

	if err != nil {
		return err
	}

	return nil
}

func (r *OpportunityRepository) GetOpportunities() ([]models.Opportunity, error) {

	query := `SELECT id,
    company_name,
    role_applied,
    application_status,
    next_step_todo,
    interview_date,
    meeting_link,
    test_link,
    comments,
    created_at,
    updated_at,
	source FROM opportunities`
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
			&opportunity.Source,
		)

		if err != nil {
			return nil, err
		}

		opportunities = append(opportunities, opportunity)
	}

	return opportunities, nil
}

func (r *OpportunityRepository) GetOpportunityByID(id int64) (*models.Opportunity, error) {

	query := `SELECT id,
    company_name,
    role_applied,
    application_status,
    next_step_todo,
    interview_date,
    meeting_link,
    test_link,
    comments,
    created_at,
    updated_at,
	source FROM opportunities WHERE id=$1`

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
		&opportunity.Source,
	)
	if err != nil {
		return nil, err
	}

	return &opportunity, nil
}

func (r *OpportunityRepository) CheckExists(companyName, roleApplied string) (int64, error) {

	query := `SELECT id FROM opportunities WHERE company_name=$1 AND role_applied=$2`

	var id int64
	err := r.db.QueryRow(query, companyName, roleApplied).Scan(&id)

	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OpportunityRepository) UpdateOpportunity(opportunity models.Opportunity, id int64) error {

	query := `UPDATE opportunities
SET
    application_status = $1,
    next_step_todo = $2,
    interview_date = $3,
    meeting_link = $4,
    test_link = $5,
    comments = $6,
	updated_at=CURRENT_TIMESTAMP
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

func (r *OpportunityRepository) SearchSimilar(embedding []float64, limit int) ([]models.Opportunity, error) {

	query := `SELECT
    id,
    company_name,
    role_applied,
    application_status,
    next_step_todo,
    interview_date,
    meeting_link,
    test_link,
    comments,
	source,
    embedding <=> $1 AS distance
FROM opportunities
WHERE embedding IS NOT NULL
ORDER BY distance
LIMIT $2`

	vector := pgvector.NewVector(extractor.ToFloat32(embedding))

	rows, err := r.db.Query(query, vector, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var opportunities []models.Opportunity
	for rows.Next() {

		var opportunity models.Opportunity
		var distance float64

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
			&opportunity.Source,
			&distance,
		)

		if err != nil {
			return nil, err
		}
		opportunities = append(opportunities, opportunity)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return opportunities, nil
}
