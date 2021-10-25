package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Employee is the structure representing a Employee.
type Employee struct {
	ID        ID        `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	Position  string    `json:"position"`
	Salary    int       `json:"salary"`
	Projects  []Project `gorm:"many2many:project_employees;" json:"projects"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EmployeeLogin struct {
	Email    string `json: "email`
	Password string `json: "password"`
}

// IsNode is Node type interface method
func (Employee) IsNode() {}

// TableName overrides the default table name...
func (Employee) TableName() string {
	return "employees"
}

//BeforeCreate -> Called before inserting record into Employees Table
func (e *Employee) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	e.ID = ID(id)
	return err
}
