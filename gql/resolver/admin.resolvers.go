package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"go-graphql-api/gql/generated"
	"go-graphql-api/middleware/auth"
	"go-graphql-api/model"
	"go-graphql-api/package/admin"
)

func (r *mutationResolver) CreateAdmin(ctx context.Context, input generated.NewAdmin) (*model.Admin, error) {
	newAdmin := model.Admin{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password,
		Role:     input.Role,
	}
	admin, err := admin.ForContext(ctx).RegisterAdmin(ctx, &newAdmin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *mutationResolver) LoginAdmin(ctx context.Context, input generated.AdminLogin) (string, error) {
	loginCredentials := model.AdminLogin{
		Email:    input.Email,
		Password: input.Password,
	}
	adminService := admin.ForContext(ctx)
	admin, err := adminService.VerifyAdmin(ctx, &loginCredentials)
	if err != nil {
		return "", err
	}
	tokenString, err := adminService.GenerateAccessToken(ctx, admin)
	if err != nil {
		return "", err
	}
	return *tokenString, nil
}

func (r *queryResolver) Admins(ctx context.Context) ([]*model.Admin, error) {
	if authSession := auth.ForContext(ctx); authSession.Role != "super-admin" {
		return nil, errors.New("Access Denied")
	}
	admins, err := admin.ForContext(ctx).GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return admins, nil
}
