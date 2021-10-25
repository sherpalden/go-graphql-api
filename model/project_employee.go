package model

import "time"

type ProjectEmployee struct {
	ProjectID  ID        `gorm:"primaryKey" json:"project_id" binding:"required"`
	EmployeeID ID        `gorm:"primaryKey" json:"employee_id" binding:"required"`
	Project    Project   `json: "project"`
	Employee   Employee  `json: "employee"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName gives table name of model
func (p ProjectEmployee) TableName() string {
	return "project_employees"
}
