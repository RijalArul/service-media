package repositories

import (
	models "service-media/models/entity"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo models.Photo) (models.Photo, error)
	FindAll() ([]models.Photo, error)
}

type PhotoRepositoryImpl struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &PhotoRepositoryImpl{DB: db}
}

func (r *PhotoRepositoryImpl) Create(photo models.Photo) (models.Photo, error) {
	err := r.DB.Create(&photo).Error
	return photo, err
}

func (r *PhotoRepositoryImpl) FindAll() ([]models.Photo, error) {
	Photos := []models.Photo{}

	err := r.DB.Find(&Photos).Error

	return Photos, err

}
