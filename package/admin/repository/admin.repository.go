package repository

import (
	"go-graphql-api/infrastructure"
	"go-graphql-api/model"
	"go-graphql-api/package/admin"

	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type adminRepository struct {
	db infrastructure.Database
}

type AdminRepositoryTarget struct {
	fx.In
	DB infrastructure.Database
}

func NewAdminRepository(target AdminRepositoryTarget) admin.Repository {
	return &adminRepository{
		db: target.DB,
	}
}

func (ar *adminRepository) GetAll() ([]*model.Admin, error) {
	var admins []*model.Admin
	return admins, ar.db.DB.Model(&model.Admin{}).Find(&admins).Error
}

func (ar *adminRepository) GetByID(id model.ID) (*model.Admin, error) {
	var admin model.Admin
	return &admin, ar.db.DB.Model(&model.Admin{}).Where("id = ?", id).Find(&admin).Error
}

func (ar *adminRepository) GetByEmail(email string) (*model.Admin, error) {
	var admin model.Admin
	err := ar.db.DB.Where("email = ?", email).First(&admin).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, err
		}
	}
	return &admin, nil
}

func (ar *adminRepository) Create(admin *model.Admin) (*model.Admin, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	admin.Password = string(hash)
	return admin, ar.db.DB.Model(&model.Admin{}).Create(admin).Error
}

func (ar *adminRepository) UpdateWithMap(id model.ID, adminMap map[string]interface{}) (*model.Admin, error) {
	var admin model.Admin
	return &admin, ar.db.DB.
		Model(&model.Admin{}).
		Where("id = ?", id).
		Updates(adminMap).
		Find(&admin).
		Error
}

func (ar *adminRepository) DeleteByID(id model.ID) error {
	return ar.db.DB.Where("id = ?", id).Delete(&model.Admin{}).Error
}
