package repository

import (
	"github.com/kittyguys/hash/api/model"
	"github.com/labstack/echo"
)

// UserRepository Define user method
type UserRepository interface {
	SignUp(u *model.User) error
	Login(t *string, n echo.Map) error
	GetUser(u *model.User, t *[]model.Tag, id string) error
	CreateTag(u *model.User, t *[]model.Tag, b map[string]interface{}) error
}
