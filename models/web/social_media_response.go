package web

import "time"

type SocialMediaCreateResponse struct {
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type SocialMediaResponse struct {
	Id             uint        `json:"id"`
	Name           string      `json:"name"`
	SocialMediaUrl string      `json:"social_media_url"`
	UserID         uint        `json:"user_id"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	User           interface{} `json:"user"`
}

type SocialMediaUpdateResponse struct {
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}
