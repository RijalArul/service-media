package web

type PhotoRequest struct {
	Title    string `json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url"`
}
