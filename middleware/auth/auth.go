package auth

import (
	"context"
	"errors"
	"go-graphql-api/constant"
	"go-graphql-api/infrastructure"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware -> structure
type AuthMiddleware struct {
	env    infrastructure.Env
	logger infrastructure.Logger
}

// NewAuthMiddleware -> creates new  autn m
func NewAuthMiddleware(
	env infrastructure.Env,
	logger infrastructure.Logger,
) AuthMiddleware {
	return AuthMiddleware{
		env:    env,
		logger: logger,
	}
}

type AuthSession struct {
	ID   string
	Role string
}

// Handle -> handles auth requests
func (m AuthMiddleware) HandleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		// Allow unauthenticated users in
		if accessToken == "" {
			c.Next()
			return
		}
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
				err := errors.New("Invalid token")
				return nil, err
			}
			return []byte(m.env.JWTSecret), nil
		})
		if err != nil {
			c.Abort()
			m.logger.Zap.Info("Invalid token")
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		id, ok := claims["id"].(string)
		if !ok {
			c.Abort()
			m.logger.Zap.Info("Claims['id'] assertion error")
		}
		role, ok := claims["role"].(string)
		if !ok {
			c.Abort()
			m.logger.Zap.Info("Claims['role'] assertion error")
		}
		authSession := AuthSession{
			ID:   id,
			Role: role,
		}
		ctx := context.WithValue(c.Request.Context(), constant.Session, &authSession)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		return
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *AuthSession {
	raw, _ := ctx.Value(constant.Session).(*AuthSession)
	return raw
}
