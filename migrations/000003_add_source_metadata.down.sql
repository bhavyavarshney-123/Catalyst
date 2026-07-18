ALTER TABLE opportunities
DROP CONSTRAINT unique_source_external_id;

ALTER TABLE opportunities
DROP COLUMN external_id;

ALTER TABLE opportunities
DROP COLUMN source;