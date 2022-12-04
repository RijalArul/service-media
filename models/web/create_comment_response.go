package web

import "time"

type CreateCommentResponse struct {
	Id        uint        `json:"id"`
	Message   string      `json:"message"`
	UserID    uint        `json:"user_id"`
	PhotoID   uint        `json:"comment_id"`
	User      interface{} `json:"user"`
	Photo     interface{} `json:"photo"`
	CreatedAt time.Time   `json:"created_at"`
}

type UpdateCommentResponse struct {
	Id        uint        `json:"id"`
	Message   string      `json:"message"`
	UserID    uint        `json:"user_id"`
	PhotoID   uint        `json:"comment_id"`
	User      interface{} `json:"user"`
	Photo     interface{} `json:"photo"`
	UpdatedAt time.Time   `json:"updated_at"`
}
