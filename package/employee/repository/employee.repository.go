package repository

import (
	"go-graphql-api/infrastructure"
	"go-graphql-api/model"
	"go-graphql-api/package/employee"

	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type employeeRepository struct {
	db infrastructure.Database
}

type EmployeeRepositoryTarget struct {
	fx.In
	DB infrastructure.Database
}

func NewEmployeeRepository(target EmployeeRepositoryTarget) employee.Repository {
	return &employeeRepository{
		db: target.DB,
	}
}

func (er *employeeRepository) GetAll() ([]*model.Employee, error) {
	var employees []*model.Employee
	return employees, er.db.DB.
		Model(&model.Employee{}).
		Preload("Projects").
		Find(&employees).Error
}

func (er *employeeRepository) GetByID(id model.ID) (*model.Employee, error) {
	var employee model.Employee
	return &employee, er.db.DB.
		Model(&model.Employee{}).
		Where("id = ?", id).
		Preload("Projects").
		Find(&employee).Error
}

func (er *employeeRepository) GetByEmail(email string) (*model.Employee, error) {
	var employee model.Employee
	err := er.db.DB.Where("email = ?", email).
		Preload("Projects").
		First(&employee).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, err
		}
	}
	return &employee, nil
}

func (er *employeeRepository) Create(employee *model.Employee) (*model.Employee, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	employee.Password = string(hash)
	return employee, er.db.DB.Model(&model.Employee{}).Create(employee).Error
}
