package repository

import (
	"go-graphql-api/infrastructure"
	"go-graphql-api/model"
	"go-graphql-api/package/project"

	"go.uber.org/fx"
)

type projectRepository struct {
	db infrastructure.Database
}

type ProjectRepositoryTarget struct {
	fx.In
	DB infrastructure.Database
}

func NewProjectRepository(target ProjectRepositoryTarget) project.Repository {
	return &projectRepository{
		db: target.DB,
	}
}

func (er *projectRepository) GetAll() ([]*model.Project, error) {
	var projects []*model.Project
	return projects, er.db.DB.Model(&model.Project{}).Find(&projects).Error
}

func (er *projectRepository) GetByID(id model.ID) (*model.Project, error) {
	var project model.Project
	return &project, er.db.DB.Model(&model.Project{}).Where("id = ?", id).Find(&project).Error
}

func (er *projectRepository) Create(project *model.Project) (*model.Project, error) {
	return project, er.db.DB.Model(&model.Project{}).Create(project).Error
}

func (er *projectRepository) AddEmployeeToProject(input *model.ProjectEmployee) (*model.Project, error) {
	err := er.db.DB.Model(&model.ProjectEmployee{}).Create(input).Error
	if err != nil {
		return nil, err
	}
	var project model.Project
	return &project, er.db.DB.Model(&model.Project{}).Where("id = ?", input.ProjectID).Find(&project).Error
}

func (er *projectRepository) GetProjectMembers(projectID model.ID, role string) ([]*model.Employee, error) {
	var projectEmployees []*model.ProjectEmployee
	err := er.db.DB.Model(&model.ProjectEmployee{}).
		Where("project_id = ? AND role = ?", projectID, role).
		Preload("Employee").
		Find(&projectEmployees).Error
	if err != nil {
		return nil, err
	}
	var employeeList []*model.Employee
	for _, val := range projectEmployees {
		employeeList = append(employeeList, &val.Employee)
	}
	return employeeList, nil
}
