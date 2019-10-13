package model

import (
	"github.com/jinzhu/gorm"
)

// Subtag TagSchema
type Subtag struct {
	gorm.Model
	Name string `json:"name"`
}

// Subtags List of Tag
type Subtags struct {
	T []Tag
}
