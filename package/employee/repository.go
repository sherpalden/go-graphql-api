package employee

import (
	"go-graphql-api/model"
)

type Repository interface {
	GetAll() ([]*model.Employee, error)
	GetByID(id model.ID) (*model.Employee, error)
	GetByEmail(email string) (*model.Employee, error)
	Create(employee *model.Employee) (*model.Employee, error)
}
