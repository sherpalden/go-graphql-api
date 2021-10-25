package employee

import (
	"context"
	"go-graphql-api/model"
)

type key string

const (
	// Employee key for employee service in each request
	Key key = "employee"
)

type Service interface {
	GetAll(ctx context.Context) ([]*model.Employee, error)
	GetByID(ctx context.Context, id model.ID) (*model.Employee, error)
	GetByEmail(ctx context.Context, email string) (*model.Employee, error)
	Create(ctx context.Context, employee *model.Employee) (*model.Employee, error)
}

// ForContext is method to get employee service from context
func ForContext(ctx context.Context) Service {
	service, ok := ctx.Value(Key).(Service)
	if !ok {
		panic("ctx passing is not contain employee service")
	}
	return service
}
