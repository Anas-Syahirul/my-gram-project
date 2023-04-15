package models

import "time"

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption  string `json:"caption" form:"caption" valid:"-"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo Url is required"`
	UserId   uint   `json:"user_id"`
	User     *User  `json:",omitempty"`
}

type CreatePhotoRequest struct {
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url" binding:"required"`
}

type CreatePhotoResponse struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserId    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type UpdatePhotoRequest struct {
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url" binding:"required"`
}

type UpdatePhotoResponse struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserId    uint       `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type GetPhotoResponse struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserId    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      *UserPhotoResponse
}

type PhotoCommentResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   uint   `json:"user_id"`
}
