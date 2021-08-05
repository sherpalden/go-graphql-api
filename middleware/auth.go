package middleware

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
		authSession := AuthSession{
			ID:   claims["id"],
			Role: claims["role"],
		}
		ctx := context.WithValue(c.Request.Context(), constant.Session, &authSession)
		c.Request = c.Request.WithContext(ctx)
		c.Writer = &authSession
		c.Next()
		return
	}
}
