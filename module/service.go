package module

import (
	admin "go-graphql-api/package/admin/service"
	employee "go-graphql-api/package/employee/service"
	project "go-graphql-api/package/project/service"

	"go.uber.org/fx"
)

// ServiceModule is Repositories fx module
var ServiceModule = fx.Provide(
	admin.NewAdminService,
	employee.NewEmployeeService,
	project.NewProjectService,
)
