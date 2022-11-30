package web

type UpdateUserRequest struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
}
