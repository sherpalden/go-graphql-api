package service

import (
	"context"

	"go-graphql-api/infrastructure"
	"go-graphql-api/model"
	"go-graphql-api/package/project"

	"go.uber.org/fx"
)

type projectService struct {
	repository project.Repository
	logger     infrastructure.Logger
	env        infrastructure.Env
}

type ProjectServiceTarget struct {
	fx.In
	Repository project.Repository
	Logger     infrastructure.Logger
	Env        infrastructure.Env
}

func NewProjectService(target ProjectServiceTarget) project.Service {
	return &projectService{
		repository: target.Repository,
		logger:     target.Logger,
		env:        target.Env,
	}
}

func (ps *projectService) GetAll(ctx context.Context) ([]*model.Project, error) {
	ps.logger.Zap.Info("------- GO_GRAPHQL_API HITTING GETALL QUERY -------")
	return ps.repository.GetAll()
}

func (ps *projectService) GetByID(ctx context.Context, id model.ID) (*model.Project, error) {
	return ps.repository.GetByID(id)
}

func (ps *projectService) Create(ctx context.Context, project *model.Project) (*model.Project, error) {
	return ps.repository.Create(project)
}

func (ps *projectService) AddEmployeeToProject(ctx context.Context, input *model.ProjectEmployee) (*model.Project, error) {
	return ps.repository.AddEmployeeToProject(input)
}

func (ps *projectService) GetProjectMembers(ctx context.Context, projectID model.ID, role string) ([]*model.Employee, error) {
	return ps.repository.GetProjectMembers(projectID, role)
}
