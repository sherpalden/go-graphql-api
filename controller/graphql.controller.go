package controller

import (
	context "context"

	"go-graphql-api/gql/generated"
	"go-graphql-api/gql/resolver"
	"go-graphql-api/infrastructure"
	"go-graphql-api/middleware"
	"go-graphql-api/package/admin"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// GraphQLController handle the graphql request, parse request to schema and return results
type GraphQLController struct {
	admin          admin.Service
	logger         infrastructure.Logger
	env            infrastructure.Env
	authMiddleware middleware.AuthMiddleware
}

// GraphQLControllerTarget is parameter object for geting all GraphQLController's dependency
type GraphQLControllerTarget struct {
	fx.In
	Admin          admin.Service
	Logger         infrastructure.Logger
	Env            infrastructure.Env
	AuthMiddleware middleware.AuthMiddleware
}

// NewGraphQLController is a constructor for GraphQLController
func NewGraphQLController(target GraphQLControllerTarget) GraphQLController {
	return GraphQLController{
		admin:          target.Admin,
		logger:         target.Logger,
		env:            target.Env,
		authMiddleware: target.AuthMiddleware,
	}
}

// GrqphQL is defining as the GraphQL handler
func (m *GraphQLController) GrqphQLHandler() gin.HandlerFunc {
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// GraphiQLHandler is defining as the GraphiQLHandler Page handler
func (m *GraphQLController) GraphiQLHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Register is function to register all controller's endpoint handler
func (m *GraphQLController) Register(r *gin.Engine) {
	r.Use(m.Middleware()).
		Use(m.authMiddleware.HandleAuth()).
		POST("/query", m.GrqphQLHandler())
	r.GET("/", m.GraphiQLHandler())
}

// Middleware for GraphQL resolver to pass services into ctx
func (m *GraphQLController) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, admin.Key, m.admin)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
