package project

import (
	"go-graphql-api/model"
)

type Repository interface {
	GetAll() ([]*model.Project, error)
	GetByID(id model.ID) (*model.Project, error)
	Create(project *model.Project) (*model.Project, error)
	AddEmployeeToProject(input *model.ProjectEmployee) (*model.Project, error)
	GetProjectMembers(projectID model.ID, role string) ([]*model.Employee, error)
}
