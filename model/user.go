package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User UserModel
type User struct {
	gorm.Model
	HashID          string        `json:"hashID"`
	DisplayName     string        `json:"displayName"`
	ProfileImageURL string        `json:"profileImageURL"`
	Email           string        `json:"email"`
	Password        string        `json:"password"`
	Tags            []Tag         `gorm:"many2many:user_tags;"`
	SocialLogin     []SocialLogin `gorm:"many2many:user_socialLogin;"`
}

// TwitterUser twitter user
type TwitterUser struct {
	ID          string     `gorm:"primary_key;not null" json:"id_str"`
	ScreenName  string     `gorm:"not null" json:"screen_name"`
	Name        string     `gorm:"not null" json:"name"`
	Image       string     `gorm:"not null" json:"profile_image_url_https"`
	Description string     `gorm:"null" json:"description"`
	IsSignedIn  bool       `gorm:"not null" json:"is_signed_in"`
	CreatedAt   time.Time  `gorm:"null" json:"create_at"`
	UpdatedAt   time.Time  `gorm:"null" json:"update_at"`
	DeletedAt   *time.Time `gorm:"null" json:"-"`
}

// Users Slice of User
var Users []User
