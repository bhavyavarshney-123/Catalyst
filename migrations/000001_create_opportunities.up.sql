CREATE TABLE opportunities (
    id BIGSERIAL PRIMARY KEY,
    company_name TEXT,
    role_applied TEXT,
    application_status TEXT,
    next_step_todo TEXT,
    interview_date TIMESTAMP,
    meeting_link TEXT,
    test_link TEXT,
    comments TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
);