package module

import (
	admin "go-graphql-api/package/admin/repository"
	employee "go-graphql-api/package/employee/repository"
	project "go-graphql-api/package/project/repository"

	"go.uber.org/fx"
)

// RepositoryModule is Repositories fx module
var RepositoryModule = fx.Provide(
	admin.NewAdminRepository,
	employee.NewEmployeeRepository,
	project.NewProjectRepository,
)
