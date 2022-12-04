package services

import (
	models "service-media/models/entity"
	"service-media/models/web"
	"service-media/repositories"
)

type SocialMediaService interface {
	Create(socmedInput web.SocialMediaRequest, userId uint) (web.SocialMediaCreateResponse, error)
	MySocialMedia(socmedia models.SocialMedia, userId uint) ([]web.SocialMediaResponse, error)
	Update(socmedInput web.SocialMediaRequest, userId uint, socialMediaId uint) (web.SocialMediaUpdateResponse, error)
}

type SocialMediaServiceImpl struct {
	SocialMediaRepository repositories.SocialMediaRepository
}

func NewSocialMediaService(socmedRepo repositories.SocialMediaRepository) SocialMediaService {
	return &SocialMediaServiceImpl{SocialMediaRepository: socmedRepo}
}

func ConvertBodyCreateSocialMediaResp(socialMedia models.SocialMedia) web.SocialMediaCreateResponse {
	return web.SocialMediaCreateResponse{
		Id:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		CreatedAt:      *socialMedia.CreatedAt,
	}
}

func ConvertBodySocialMediaResp(socialMedia models.SocialMedia) web.SocialMediaResponse {
	return web.SocialMediaResponse{
		Id:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		User:           convertBodyPhotoUser(*socialMedia.User),
	}
}

func ConvertBodyUpdateSocialMediaResp(socialMedia models.SocialMedia) web.SocialMediaUpdateResponse {
	return web.SocialMediaUpdateResponse{
		Id:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		UpdatedAt:      *socialMedia.UpdatedAt,
	}
}

func (s *SocialMediaServiceImpl) Create(socmedInput web.SocialMediaRequest, userId uint) (web.SocialMediaCreateResponse, error) {
	socmed := models.SocialMedia{
		Name:           socmedInput.Name,
		SocialMediaUrl: socmedInput.SocialMediaUrl,
		UserID:         userId,
	}

	newSocmed, err := s.SocialMediaRepository.Create(socmed)
	return ConvertBodyCreateSocialMediaResp(newSocmed), err
}

func (s *SocialMediaServiceImpl) MySocialMedia(socmed models.SocialMedia, userId uint) ([]web.SocialMediaResponse, error) {
	socmed.UserID = userId
	socmedResp := []web.SocialMediaResponse{}
	getSocmed, err := s.SocialMediaRepository.GetSocialMedia(socmed)

	for i := 0; i < len(getSocmed); i++ {
		socialmedia := ConvertBodySocialMediaResp(getSocmed[i])
		socmedResp = append(socmedResp, socialmedia)
	}

	return socmedResp, err

}

func (s *SocialMediaServiceImpl) Update(socmedInput web.SocialMediaRequest, userId uint, socialMediaId uint) (web.SocialMediaUpdateResponse, error) {
	socialMedia := models.SocialMedia{
		Name:           socmedInput.Name,
		SocialMediaUrl: socmedInput.SocialMediaUrl,
		UserID:         userId,
	}
	updateSocmed, err := s.SocialMediaRepository.Update(socialMedia, socialMediaId)

	return ConvertBodyUpdateSocialMediaResp(updateSocmed), err
}
