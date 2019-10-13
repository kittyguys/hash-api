package model

import (
	"github.com/jinzhu/gorm"
)

// Tag TagSchema
type Tag struct {
	gorm.Model
	Name    string   `json:"name"`
	Users   []User   `gorm:"many2many:user_tags;"`
	Subtags []Subtag `gorm:"many2many:tag_subtags;"`
}

// Tags List of Tag
type Tags struct {
	T []Tag
}

// Create リクエストBODY
type Create struct {
	Tag    string `json:"tag"`
	HashID string `json:"hashID"`
}
