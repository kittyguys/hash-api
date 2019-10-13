package model

import (
	"github.com/jinzhu/gorm"
)

// SocialLogin SocialLogin
type SocialLogin struct {
	gorm.Model
	Name string `json:"name"`
}
