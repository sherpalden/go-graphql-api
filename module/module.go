package module

import (
	"context"
	"go-graphql-api/controller"
	"go-graphql-api/infrastructure"
	"go-graphql-api/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Module is registry for all module using in application
// will process by fx
var Module = fx.Options(
	ServiceModule,
	RepositoryModule,
	InfrastructureModule,
	fx.Provide(
		controller.NewGraphQLController,
		middleware.NewAuthMiddleware,
	),
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	env infrastructure.Env,
	logger infrastructure.Logger,
	database infrastructure.Database,
	controller controller.GraphQLController,
) {
	conn, _ := database.DB.DB()
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("----------------------------")
			logger.Zap.Info("------- GO_GRAPHQL_API -------")
			logger.Zap.Info("----------------------------")

			conn.SetMaxOpenConns(10)

			go func() {
				httpRouter := gin.Default()
				controller.Register(httpRouter)
				httpRouter.Run(":" + env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping Application")
			conn.Close()
			return nil
		},
	})
}
