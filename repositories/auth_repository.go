package repositories

import (
	"depoguna-api/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user *models.User) error
	Login(email string) (*models.User, error)
	IsDuplicateEmail(email string) bool
}

type authReposiitory struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authReposiitory{
		DB: db,
	}
}

func (r *authReposiitory) Register(user *models.User) error {
	return r.DB.Table("users").Debug().Create(&user).Error
}

func (r *authReposiitory) Login(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authReposiitory) IsDuplicateEmail(email string) bool {
	var user models.User
	if err := r.DB.Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return true
}
