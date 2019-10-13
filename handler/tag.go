package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/kittyguys/hash/api/interfaces"
	"github.com/kittyguys/hash/api/model"
	"github.com/kittyguys/hash/api/repository"
	"github.com/labstack/echo"
)

// NewTagHandler Initialize user repository
func NewTagHandler(conn *gorm.DB) *TagHandler {
	return &TagHandler{
		repo: interfaces.NewTagRepo(conn),
	}
}

// TagHandler Handler with DB
type TagHandler struct {
	repo repository.TagRepository
}

// GetUsers ユーザーにタグを追加する
func (h *TagHandler) GetUsers(c echo.Context) (err error) {
	var tag model.Tag
	var users []model.User
	name := c.Param("name")

	h.repo.GetUsers(&tag, &users, name)

	data := map[string]interface{}{"users": users}

	return c.JSON(http.StatusCreated, data)
}
