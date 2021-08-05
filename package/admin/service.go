package admin

import (
	"context"
	"go-graphql-api/model"
)

type key string

const (
	// Admin key for admin service in each request
	Key key = "admin"
)

type Service interface {
	GetAll(ctx context.Context) ([]*model.Admin, error)
	GetByID(ctx context.Context, id model.ID) (*model.Admin, error)
	GetByEmail(ctx context.Context, email string) (*model.Admin, error)
	Create(ctx context.Context, admin *model.Admin) (*model.Admin, error)
	UpdateWithMap(ctx context.Context, id model.ID, adminMap map[string]interface{}) (*model.Admin, error)
	DeleteByID(ctx context.Context, id model.ID) error

	RegisterAdmin(ctx context.Context, admin *model.Admin) (*model.Admin, error)
	VerifyAdmin(ctx context.Context, adminLogin *model.AdminLogin) (*model.Admin, error)
	GenerateAccessToken(ctx context.Context, admin *model.Admin) (*string, error)
}

// ForContext is method to get admin service from context
func ForContext(ctx context.Context) Service {
	service, ok := ctx.Value(Key).(Service)
	if !ok {
		panic("ctx passing is not contain admin service")
	}
	return service
}
