package repositories

import (
	models "service-media/models/entity"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo models.Photo) (models.Photo, error)
	FindAll() ([]models.Photo, error)
	FindPhotosByUser(userId uint) ([]models.Photo, error)
	FindPhotoById(id uint) (models.Photo, error)
	UpdatePhoto(photo models.Photo, photoId uint) (models.Photo, error)
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

func (r *PhotoRepositoryImpl) FindPhotosByUser(userId uint) ([]models.Photo, error) {
	Photos := []models.Photo{}

	err := r.DB.Find(&Photos, "user_id = ?", userId).Error

	return Photos, err
}

func (r *PhotoRepositoryImpl) FindPhotoById(id uint) (models.Photo, error) {
	photo := models.Photo{}
	err := r.DB.Model(&photo).First(&photo, id).Error
	return photo, err
}

func (r *PhotoRepositoryImpl) UpdatePhoto(photo models.Photo, photoId uint) (models.Photo, error) {
	err := r.DB.Model(&photo).Where("id = ?", photoId).Updates(&photo).First(&photo).Error
	return photo, err
}
