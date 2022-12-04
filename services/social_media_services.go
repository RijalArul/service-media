package services

import (
	models "service-media/models/entity"
	"service-media/models/web"
	"service-media/repositories"
)

type SocialMediaService interface {
	Create(socmedInput web.SocialMediaRequest, userId uint) (web.SocialMediaResponse, error)
	MySocialMedia(socmedia models.SocialMedia, userId uint) ([]web.SocialMediaResponse, error)
}

type SocialMediaServiceImpl struct {
	SocialMediaRepository repositories.SocialMediaRepository
}

func NewSocialMediaService(socmedRepo repositories.SocialMediaRepository) SocialMediaService {
	return &SocialMediaServiceImpl{SocialMediaRepository: socmedRepo}
}

func ConvertBodySocialMediaResp(socialMedia models.SocialMedia) web.SocialMediaResponse {
	return web.SocialMediaResponse{
		Id:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		CreatedAt:      *socialMedia.CreatedAt,
	}
}

func (s *SocialMediaServiceImpl) Create(socmedInput web.SocialMediaRequest, userId uint) (web.SocialMediaResponse, error) {
	socmed := models.SocialMedia{
		Name:           socmedInput.Name,
		SocialMediaUrl: socmedInput.SocialMediaUrl,
		UserID:         userId,
	}

	newSocmed, err := s.SocialMediaRepository.Create(socmed)
	return ConvertBodySocialMediaResp(newSocmed), err
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
