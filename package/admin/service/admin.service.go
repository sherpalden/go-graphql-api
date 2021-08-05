package service

import (
	"context"
	"errors"
	"time"

	"go-graphql-api/infrastructure"
	"go-graphql-api/model"
	"go-graphql-api/package/admin"

	jwt "github.com/dgrijalva/jwt-go"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
)

type adminService struct {
	repository admin.Repository
	logger     infrastructure.Logger
	env        infrastructure.Env
}

type AdminServiceTarget struct {
	fx.In
	Repository admin.Repository
	Logger     infrastructure.Logger
	Env        infrastructure.Env
}

func NewAdminService(target AdminServiceTarget) admin.Service {
	return &adminService{
		repository: target.Repository,
		logger:     target.Logger,
		env:        target.Env,
	}
}

func (as *adminService) RegisterAdmin(ctx context.Context, admin *model.Admin) (*model.Admin, error) {
	oldAdmin, err := as.repository.GetByEmail(admin.Email)
	if oldAdmin != nil {
		return nil, errors.New("Admin with this email exits already")
	}
	admin, err = as.repository.Create(admin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (as *adminService) VerifyAdmin(ctx context.Context, adminLogin *model.AdminLogin) (*model.Admin, error) {
	admin, err := as.repository.GetByEmail(adminLogin.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(adminLogin.Password))
	if err != nil {
		return nil, err
	}
	return admin, err
}

func (as *adminService) GetAccessToken(ctx context.Context, admin *model.Admin) (*string, error) {
	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	// Set some claims
	claims["id"] = admin.ID
	claims["role"] = admin.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(as.env.JWTSecret))
	if err != nil {
		return nil, err
	}
	return &tokenString, err
}

func (as *adminService) GetAll(ctx context.Context) ([]*model.Admin, error) {
	as.logger.Zap.Info("------- GO_GRAPHQL_API HITTING GETALL QUERY -------")
	return as.repository.GetAll()
}

func (as *adminService) GetByID(ctx context.Context, id model.ID) (*model.Admin, error) {
	return as.repository.GetByID(id)
}

func (as *adminService) GetByEmail(ctx context.Context, email string) (*model.Admin, error) {
	return as.repository.GetByEmail(email)
}

func (as *adminService) Create(ctx context.Context, admin *model.Admin) (*model.Admin, error) {
	return as.repository.Create(admin)
}

func (as *adminService) UpdateWithMap(ctx context.Context, id model.ID, adminMap map[string]interface{}) (*model.Admin, error) {
	return as.repository.UpdateWithMap(id, adminMap)
}

func (as *adminService) DeleteByID(ctx context.Context, id model.ID) error {
	return as.repository.DeleteByID(id)
}
