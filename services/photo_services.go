package services

import (
	models "service-media/models/entity"
	"service-media/models/web"
	"service-media/repositories"
)

type PhotoService interface {
	Create(PhotoInput web.CreaatePhotoRequest, userId uint) (web.CreatePhotoResponse, error)
}

type PhotoServiceImpl struct {
	PhotoRepository repositories.PhotoRepository
}

func NewPhotoService(photoRepository *repositories.PhotoRepository) PhotoService {
	return &PhotoServiceImpl{PhotoRepository: *photoRepository}
}

func convertBodyPhotoResponse(photo models.Photo) web.CreatePhotoResponse {
	return web.CreatePhotoResponse{
		Id:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID:   photo.UserID,
	}
}

func (s *PhotoServiceImpl) Create(PhotoInput web.CreaatePhotoRequest, userId uint) (web.CreatePhotoResponse, error) {
	photo := models.Photo{
		Title:    PhotoInput.Title,
		Caption:  PhotoInput.Caption,
		PhotoUrl: PhotoInput.PhotoUrl,
		UserID:   userId,
	}
	newPhoto, err := s.PhotoRepository.Create(photo)
	return convertBodyPhotoResponse(newPhoto), err
}
