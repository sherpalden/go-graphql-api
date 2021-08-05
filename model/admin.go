package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Admin is the structure representing a admin.
type Admin struct {
	ID        ID        `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminLogin struct {
	Email    string `json: "email`
	Password string `json: "password"`
}

// IsNode is Node type interface method
func (Admin) IsNode() {}

// TableName overrides the default table name...
func (Admin) TableName() string {
	return "admins"
}

//BeforeCreate -> Called before inserting record into admins Table
func (a *Admin) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	a.ID = ID(id)
	return err
}
