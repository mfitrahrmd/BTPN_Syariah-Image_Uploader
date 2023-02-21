package app

type Photo struct {
	ID       uint   `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Caption  string `json:"caption,omitempty"`
	PhotoUrl string `json:"photoUrl,omitempty"`
	UserID   uint   `json:"userID,omitempty"`
}

type InsertPhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption" binding:"required"`
	PhotoUrl string `json:"photoUrl" binding:"required,url"`
}

type InsertPhotoResponse struct {
	ID       uint   `json:"id,omitempty"`
	Title    string `json:"title,omitempty" binding:"required"`
	Caption  string `json:"caption,omitempty" binding:"required"`
	PhotoUrl string `json:"photoUrl,omitempty" binding:"required,url"`
}

type FindAllPhotosResponse struct {
	Photos []Photo `json:"photos"`
}

type UpdatePhotoRequest struct {
	Title    string `json:"title,omitempty" binding:"required"`
	Caption  string `json:"caption,omitempty" binding:"required"`
	PhotoUrl string `json:"photoUrl,omitempty" binding:"required,url"`
}

type UpdatePhotoResponse struct {
	ID       uint   `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Caption  string `json:"caption,omitempty"`
	PhotoUrl string `json:"photoUrl,omitempty"`
}

type DeletePhotoResponse struct {
	ID uint `json:"id,omitempty"`
}
