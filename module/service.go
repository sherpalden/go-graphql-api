package module

import (
	admin "go-graphql-api/package/admin/service"

	"go.uber.org/fx"
)

// ServiceModule is Repositories fx module
var ServiceModule = fx.Provide(
	admin.NewAdminService,
)
