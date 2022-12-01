package services

import (
	models "service-media/models/entity"
	"service-media/models/web"
	"service-media/repositories"
)

type PhotoService interface {
	Create(PhotoInput web.CreaatePhotoRequest, userId uint) (web.CreatePhotoResponse, error)
	GetAllPhotos() ([]web.CreatePhotoResponse, error)
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

func (s *PhotoServiceImpl) GetAllPhotos() ([]web.CreatePhotoResponse, error) {
	photos, err := s.PhotoRepository.FindAll()

	var photoResp []web.CreatePhotoResponse
	for i := 0; i < len(photos); i++ {
		photoResp = []web.CreatePhotoResponse{
			{
				Id:       photos[i].ID,
				Title:    photos[i].Title,
				Caption:  photos[i].Caption,
				PhotoUrl: photos[i].PhotoUrl,
				UserID:   photos[i].UserID,
			},
		}
	}

	return photoResp, err

}
