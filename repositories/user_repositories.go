package repositories

import (
	models "service-media/models/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Create(user models.User) (models.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) FindByEmail(email string) (models.User, error) {
	user := models.User{}
	err := r.DB.Model(&user).Where("email = ?", email).First(&user).Error

	return user, err
}
