CREATE EXTENSION IF NOT EXISTS vector;

ALTER TABLE opportunities
ADD COLUMN embedding vector(1536);