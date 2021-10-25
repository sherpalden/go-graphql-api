CREATE TABLE IF NOT EXISTS projects (
  id uuid,
  name VARCHAR(100) NOT NULL,
  project_owner JSONB, 
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);
