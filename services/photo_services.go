package services

import (
	models "service-media/models/entity"
	"service-media/models/web"
	"service-media/repositories"
)

type PhotoService interface {
	Create(photoInput web.PhotoRequest, userId uint) (web.CreatePhotoResponse, error)
	GetAllPhotos() ([]web.CreatePhotoResponse, error)
	GetPhotosByUser(userId uint) ([]web.CreatePhotoResponse, error)
	UpdatePhoto(photoInput web.PhotoRequest, photoId uint) (web.UpdatePhotoResponse, error)
	DeletePhoto(photoId uint) error
}

type PhotoServiceImpl struct {
	PhotoRepository repositories.PhotoRepository
}

func NewPhotoService(photoRepository repositories.PhotoRepository) PhotoService {
	return &PhotoServiceImpl{PhotoRepository: photoRepository}
}

func convertBodyPhotoResponse(photo models.Photo) web.CreatePhotoResponse {
	return web.CreatePhotoResponse{
		Id:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		User:      convertBodyPhotoUser(*photo.User),
		CreatedAt: *photo.CreatedAt,
	}
}

func convertBodyUpdatePhotoResponse(photo models.Photo) web.UpdatePhotoResponse {
	return web.UpdatePhotoResponse{
		Id:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID:   photo.UserID,
	}
}

func convertBodyAssociatePhoto(photo models.Photo) web.CreatePhotoAssociateResp {
	return web.CreatePhotoAssociateResp{
		Id:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID:   photo.UserID,
	}
}

func (s *PhotoServiceImpl) Create(photoInput web.PhotoRequest, userId uint) (web.CreatePhotoResponse, error) {
	photo := models.Photo{
		Title:    photoInput.Title,
		Caption:  photoInput.Caption,
		PhotoUrl: photoInput.PhotoUrl,
		UserID:   userId,
	}
	newPhoto, err := s.PhotoRepository.Create(photo)
	return convertBodyPhotoResponse(newPhoto), err
}

func (s *PhotoServiceImpl) GetAllPhotos() ([]web.CreatePhotoResponse, error) {
	photos, err := s.PhotoRepository.FindAll()

	photoResp := []web.CreatePhotoResponse{}
	for i := 0; i < len(photos); i++ {
		photo := convertBodyPhotoResponse(photos[i])
		photoResp = append(photoResp, photo)
	}

	return photoResp, err
}

func (s *PhotoServiceImpl) GetPhotosByUser(userId uint) ([]web.CreatePhotoResponse, error) {
	photos, err := s.PhotoRepository.FindPhotosByUser(userId)
	photoResp := []web.CreatePhotoResponse{}
	for i := 0; i < len(photos); i++ {
		photo := convertBodyPhotoResponse(photos[i])
		photoResp = append(photoResp, photo)
	}

	return photoResp, err
}

func (s *PhotoServiceImpl) UpdatePhoto(photoInput web.PhotoRequest, photoId uint) (web.UpdatePhotoResponse, error) {
	photo := models.Photo{
		Title:    photoInput.Title,
		Caption:  photoInput.Caption,
		PhotoUrl: photoInput.PhotoUrl,
	}

	updatePhoto, err := s.PhotoRepository.UpdatePhoto(photo, photoId)
	return convertBodyUpdatePhotoResponse(updatePhoto), err
}

func (s *PhotoServiceImpl) DeletePhoto(photoId uint) error {
	err := s.PhotoRepository.DeletePhoto(photoId)
	return err
}
