package service

import (
	"context"
	"errors"

	"go-graphql-api/infrastructure"
	"go-graphql-api/model"
	"go-graphql-api/package/employee"

	"go.uber.org/fx"
)

type employeeService struct {
	repository employee.Repository
	logger     infrastructure.Logger
	env        infrastructure.Env
}

type EmployeeServiceTarget struct {
	fx.In
	Repository employee.Repository
	Logger     infrastructure.Logger
	Env        infrastructure.Env
}

func NewEmployeeService(target EmployeeServiceTarget) employee.Service {
	return &employeeService{
		repository: target.Repository,
		logger:     target.Logger,
		env:        target.Env,
	}
}

func (es *employeeService) GetAll(ctx context.Context) ([]*model.Employee, error) {
	es.logger.Zap.Info("------- GO_GRAPHQL_API HITTING GETALL QUERY -------")
	return es.repository.GetAll()
}

func (es *employeeService) GetByID(ctx context.Context, id model.ID) (*model.Employee, error) {
	return es.repository.GetByID(id)
}

func (es *employeeService) GetByEmail(ctx context.Context, email string) (*model.Employee, error) {
	return es.repository.GetByEmail(email)
}

func (es *employeeService) Create(ctx context.Context, employee *model.Employee) (*model.Employee, error) {
	oldEmployee, _ := es.repository.GetByEmail(employee.Email)
	if oldEmployee != nil {
		return nil, errors.New("Employee with this email exits already")
	}
	return es.repository.Create(employee)
}
