package models

import "time"

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" binding:"required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" binding:"required"`
	UserId         uint   `json:"user_id"`
	User           *User
}

type CreateSocialMediaResponse struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserId         uint       `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
}

type GetSocialMediaResponse struct {
	ID             	uint       	`json:"id"`
	Name           	string     	`json:"name"`
	SocialMediaUrl 	string     	`json:"social_media_url"`
	UserId         	uint       	`json:"user_id"`
	UpdatedAt 		*time.Time 	`json:"updated_at"`
	CreatedAt      	*time.Time 	`json:"created_at"`
	User           	*UserSocialMediaResponse
}

type UpdateSocialMediaResponse struct {
	ID             	uint       `json:"id"`
	Name           	string     `json:"name"`
	SocialMediaUrl 	string     `json:"social_media_url"`
	UserId         	uint       `json:"user_id"`
	UpdatedAt 		*time.Time 	`json:"updated_at"`
}