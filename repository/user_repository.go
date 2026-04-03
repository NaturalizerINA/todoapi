package repository

import (
	"todoapi/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) FindByEmail(email string) (models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (r *userRepositoryImpl) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryImpl) Update(user *models.User) error {
	return r.db.Save(user).Error
}
