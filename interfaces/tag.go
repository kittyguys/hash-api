package interfaces

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/kittyguys/hash/api/model"
	"github.com/kittyguys/hash/api/repository"
	"github.com/labstack/echo"
)

// NewTagRepo Initialize user repository
func NewTagRepo(conn *gorm.DB) repository.TagRepository {
	return &TagRepository{
		Conn: conn,
	}
}

// TagRepository Handler with DB
type TagRepository struct {
	Conn *gorm.DB
}

// GetUsers GetUsers
func (h *TagRepository) GetUsers(t *model.Tag, u *[]model.User, n string) error {
	var users []model.User
	if err := h.Conn.First(&t, model.Tag{Name: n}).Error; err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Not found"}
	}
	h.Conn.Preload("Tags").Find(&users)
	*u = users
	return nil
}
