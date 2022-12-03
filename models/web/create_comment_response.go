package web

import "time"

type CreateCommentResponse struct {
	Id        uint      `json:"id"`
	Message   string    `json:"message"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"comment_id"`
	CreatedAt time.Time `json:"created_at"`
}
