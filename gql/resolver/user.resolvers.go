package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-graphql-api/gql/generated"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input generated.NewUser) (*generated.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*generated.User, error) {
	panic(fmt.Errorf("not implemented"))
}
