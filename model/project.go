package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//ProjectOwner type
type ProjectOwner struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Project Struct
type Project struct {
	ID           ID           `json:"id"`
	Name         string       `json:"name"`
	ProjectOwner ProjectOwner `gorm:"type:JSON;" json:"project_owner"`
	Employees    []Employee   `gorm:"many2many:project_employees;" json:"employees"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

// TableName func returns the table name
func (p *Project) TableName() string {
	return "projects"
}

//BeforeCreate -> Called before inserting record into Project Table
func (p *Project) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	p.ID = ID(id)
	return err
}

//ProjectOwner data type
func (ProjectOwner) GormDataType() string {
	return "json"
}

// ProjectOwner value to DB
func (po ProjectOwner) Value() (driver.Value, error) {
	return json.Marshal(po)
}

//ProjectOwner value from DB
func (po *ProjectOwner) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var result ProjectOwner
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value"))
	}
	*po = ProjectOwner(result)
	return nil
}
