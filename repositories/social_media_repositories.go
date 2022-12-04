package repositories

import (
	models "service-media/models/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialMediaRepository interface {
	Create(socialMedia models.SocialMedia) (models.SocialMedia, error)
	GetSocialMedia(socialMedia models.SocialMedia) ([]models.SocialMedia, error)
	Update(socialMedia models.SocialMedia, socialMediaId uint) (models.SocialMedia, error)
}

type SocialMediaRepositoryImpl struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &SocialMediaRepositoryImpl{DB: db}
}

func (r *SocialMediaRepositoryImpl) Create(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	err := r.DB.Preload(clause.Associations).Create(&socialMedia).Error
	return socialMedia, err
}

func (r *SocialMediaRepositoryImpl) GetSocialMedia(socialMedia models.SocialMedia) ([]models.SocialMedia, error) {
	var socmedias []models.SocialMedia
	err := r.DB.Preload(clause.Associations).Find(&socmedias, "user_id = ?", socialMedia.UserID).Error
	return socmedias, err
}

func (r *SocialMediaRepositoryImpl) Update(socialMedia models.SocialMedia, socialMediaId uint) (models.SocialMedia, error) {
	err := r.DB.Preload(clause.Associations).Where("id = ?", socialMediaId).Updates(socialMedia).First(&socialMedia).Error
	return socialMedia, err
}
