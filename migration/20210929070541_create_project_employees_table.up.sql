CREATE TABLE project_employees (
  project_id    uuid REFERENCES projects (id),
  employee_id uuid REFERENCES employees (id),
  role VARCHAR(100) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT project_employees_pkey PRIMARY KEY (project_id, employee_id)
);