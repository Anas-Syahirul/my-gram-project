package models

import "time"

type Comment struct {
	GormModel
	UserId  uint   `json:"user_id"`
	PhotoId uint   `json:"photo_id"`
	Message string `gorm:"not null" json:"message" form:"message" binding:"required"`
	User    *User
	Photo   *Photo
}

type CreateCommentRequest struct {
	PhotoId uint   `json:"photo_id" form:"photo_id"`
	Message string `json:"message" form:"message" binding:"required"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" form:"message" binding:"required"`
}

type GetCommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	UserId    uint       `json:"user_id"`
	PhotoId   uint       `json:"photo_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      *UserCommentResponse
	Photo     *PhotoCommentResponse
}

type CreateCommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	UserId    uint       `json:"user_id"`
	PhotoId   uint       `json:"photo_id"`
	CreatedAt *time.Time `json:"updated_at"`
}

type UpdateCommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	UserId    uint       `json:"user_id"`
	PhotoId   uint       `json:"photo_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}
