package module

import (
	"go-graphql-api/infrastructure"

	"go.uber.org/fx"
)

// Module exported for initializing application
var InfrastructureModule = fx.Options(
	fx.Provide(infrastructure.NewEnv),
	fx.Provide(infrastructure.NewDatabase),
	fx.Provide(infrastructure.NewLogger),
)
