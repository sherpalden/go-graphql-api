package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-graphql-api/gql/generated"
	"go-graphql-api/model"
	"go-graphql-api/package/project"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input generated.NewProject) (*model.Project, error) {
	projectOwner := model.ProjectOwner{
		Name:  input.Owner.Name,
		Email: input.Owner.Email,
		Phone: input.Owner.Phone,
	}
	newProject := model.Project{
		Name:         input.Name,
		ProjectOwner: projectOwner,
	}
	respProject, err := project.ForContext(ctx).Create(ctx, &newProject)
	if err != nil {
		return nil, err
	}
	return respProject, nil
}

func (r *mutationResolver) AddEmployeeToProject(ctx context.Context, input generated.EmpToProjectInput) (*model.Project, error) {
	projectID, err := model.StringToID(input.ProjectID)
	if err != nil {
		return nil, err
	}
	employeeID, err := model.StringToID(input.EmployeeID)
	if err != nil {
		return nil, err
	}
	projectEmployee := model.ProjectEmployee{
		ProjectID:  projectID,
		EmployeeID: employeeID,
		Role:       input.Role.String(),
	}
	respProject, err := project.ForContext(ctx).AddEmployeeToProject(ctx, &projectEmployee)
	if err != nil {
		return nil, err
	}
	return respProject, nil
}

func (r *projectResolver) Name(ctx context.Context, obj *model.Project) (string, error) {
	return obj.Name, nil
}

func (r *projectResolver) Owner(ctx context.Context, obj *model.Project) (*generated.ProjectOwner, error) {
	projectOwner := generated.ProjectOwner{
		Name:  obj.ProjectOwner.Name,
		Email: obj.ProjectOwner.Email,
		Phone: obj.ProjectOwner.Phone,
	}
	return &projectOwner, nil
}

func (r *projectResolver) Manager(ctx context.Context, obj *model.Project) (*model.Employee, error) {
	empList, err := project.ForContext(ctx).GetProjectMembers(ctx, obj.ID, "MANAGER")
	if err != nil {
		return nil, err
	}
	if len(empList) < 1 {
		return nil, nil
	}
	return empList[0], nil
}

func (r *projectResolver) Developers(ctx context.Context, obj *model.Project) ([]*model.Employee, error) {
	empList, err := project.ForContext(ctx).GetProjectMembers(ctx, obj.ID, "DEVELOPER")
	if err != nil {
		return nil, err
	}
	return empList, nil
}

func (r *projectResolver) TeamLead(ctx context.Context, obj *model.Project) (*model.Employee, error) {
	empList, err := project.ForContext(ctx).GetProjectMembers(ctx, obj.ID, "TEAM_LEAD")
	if err != nil {
		return nil, err
	}
	if len(empList) < 1 {
		return nil, nil
	}
	return empList[0], nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	return project.ForContext(ctx).GetAll(ctx)
}

func (r *queryResolver) Project(ctx context.Context, input string) (*model.Project, error) {
	projectID, err := model.StringToID(input)
	if err != nil {
		return nil, err
	}
	respProject, err := project.ForContext(ctx).GetByID(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return respProject, nil
}

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

type projectResolver struct{ *Resolver }
