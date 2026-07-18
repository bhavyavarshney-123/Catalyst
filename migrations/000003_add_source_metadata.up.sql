ALTER TABLE opportunities
ADD COLUMN source TEXT NOT NULL DEFAULT 'gmail',
ADD COLUMN external_id TEXT NOT NULL;

ALTER TABLE opportunities
ADD CONSTRAINT unique_source_external_id
UNIQUE (source, external_id);