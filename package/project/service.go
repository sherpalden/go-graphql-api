package project

import (
	"context"
	"go-graphql-api/model"
)

type key string

const (
	// Project key for project service in each request
	Key key = "project"
)

type Service interface {
	GetAll(ctx context.Context) ([]*model.Project, error)
	GetByID(ctx context.Context, id model.ID) (*model.Project, error)
	Create(ctx context.Context, project *model.Project) (*model.Project, error)
	AddEmployeeToProject(ctx context.Context, input *model.ProjectEmployee) (*model.Project, error)
	GetProjectMembers(ctx context.Context, projectID model.ID, role string) ([]*model.Employee, error)
}

// ForContext is method to get project service from context
func ForContext(ctx context.Context) Service {
	service, ok := ctx.Value(Key).(Service)
	if !ok {
		panic("ctx passing is not contain project service")
	}
	return service
}
