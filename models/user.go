package models

type User struct {
	GormModel
	Username        string        `gorm:"not null;uniqueIndex" json:"username" form:"username" binding:"required"`
	Email           string        `gorm:"not null;uniqueIndex" json:"email" binding:"required,email"`
	Password        string        `gorm:"not null" json:"password" form:"password" binding:"required,min=6"`
	Age             int           `gorm:"not null" json:"age" form:"age" binding:"required,min=9"`
	ProfileImageUrl string        `json:"profile_image_url" form:"profile_image_url"`
	Photos          []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Comments        []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	SocialMedias    []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

type UserSocialMediaResponse struct {
	ID              uint   `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type UserPhotoResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserCommentResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}