package repository

import "github.com/kittyguys/hash/api/model"

// TagRepository Define user method
type TagRepository interface {
	GetUsers(t *model.Tag, u *[]model.User, n string) error
}
