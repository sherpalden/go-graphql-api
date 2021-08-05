package module

import (
	admin "go-graphql-api/package/admin/repository"

	"go.uber.org/fx"
)

// RepositoryModule is Repositories fx module
var RepositoryModule = fx.Provide(
	admin.NewAdminRepository,
)
