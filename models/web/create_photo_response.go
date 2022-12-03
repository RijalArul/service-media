package web

import "time"

type CreatePhotoResponse struct {
	Id        uint        `json:"id"`
	Title     string      `json:"title"`
	Caption   string      `json:"caption"`
	PhotoUrl  string      `json:"photo_url"`
	UserID    uint        `json:"user_id"`
	User      interface{} `json:"user"`
	CreatedAt time.Time   `json:"created_at"`
}

type CreatePhotoUserResp struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
