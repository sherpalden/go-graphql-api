package admin

import (
	"go-graphql-api/model"
)

type Repository interface {
	GetAll() ([]*model.Admin, error)
	GetByID(id model.ID) (*model.Admin, error)
	GetByEmail(email string) (*model.Admin, error)
	Create(admin *model.Admin) (*model.Admin, error)
	UpdateWithMap(id model.ID, adminMap map[string]interface{}) (*model.Admin, error)
	DeleteByID(id model.ID) error
}
