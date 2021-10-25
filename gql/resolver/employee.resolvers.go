package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-graphql-api/gql/generated"
	"go-graphql-api/model"
	"go-graphql-api/package/employee"
)

func (r *employeeResolver) Position(ctx context.Context, obj *model.Employee) (generated.Position, error) {
	var position generated.Position
	err := position.UnmarshalGQL(obj.Position)
	return position, err
}

func (r *mutationResolver) CreateEmployee(ctx context.Context, input generated.NewEmployee) (*model.Employee, error) {
	newEmployee := model.Employee{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password,
		Salary:   input.Salary,
		Position: input.Position.String(),
	}
	employee, err := employee.ForContext(ctx).Create(ctx, &newEmployee)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *queryResolver) Employees(ctx context.Context) ([]*model.Employee, error) {
	return employee.ForContext(ctx).GetAll(ctx)
}

func (r *queryResolver) Employee(ctx context.Context, input string) (*model.Employee, error) {
	var respEmployee *model.Employee
	employeeID, err := model.StringToID(input)
	if err != nil {
		respEmployee, err = employee.ForContext(ctx).GetByEmail(ctx, input)
		if err != nil {
			return nil, err
		}
	} else {
		respEmployee, err = employee.ForContext(ctx).GetByID(ctx, employeeID)
		if err != nil {
			return nil, err
		}
	}
	return respEmployee, nil
}

// Employee returns generated.EmployeeResolver implementation.
func (r *Resolver) Employee() generated.EmployeeResolver { return &employeeResolver{r} }

type employeeResolver struct{ *Resolver }
