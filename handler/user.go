package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kittyguys/hash/api/interfaces"
	"github.com/kittyguys/hash/api/model"
	"github.com/kittyguys/hash/api/repository"

	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
)

// NewUserHandler Initialize user repository
func NewUserHandler(conn *gorm.DB) *UserHandler {
	return &UserHandler{
		repo: interfaces.NewUserRepo(conn),
	}
}

// UserHandler Handler with DB
type UserHandler struct {
	repo repository.UserRepository
}

// SignUp sign up
func (h *UserHandler) SignUp(c echo.Context) (err error) {
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	if err := h.repo.SignUp(u); err != nil {
		return err
	}

	if u.Email != "" || u.Password != "" {

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["admin"] = true
		claims["hashID"] = u.HashID
		claims["displayName"] = u.DisplayName
		claims["iat"] = time.Now()
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		tokenString, _ := token.SignedString([]byte(Key))
		data := map[string]interface{}{"token": tokenString}

		return c.JSON(http.StatusCreated, data)
	}

	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
}

// Login log in
func (h *UserHandler) Login(c echo.Context) (err error) {
	var body echo.Map
	var token string
	if err = c.Bind(&body); err != nil {
		return
	}

	h.repo.Login(&token, body)

	if token != "" {
		data := map[string]interface{}{"token": token}

		return c.JSON(http.StatusCreated, data)
	}

	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
}

// GetUser for getting user info by ID
func (h *UserHandler) GetUser(c echo.Context) (err error) {
	var u model.User
	var tags []model.Tag
	id := c.Param("id")

	h.repo.GetUser(&u, &tags, id)

	data := map[string]interface{}{"hashID": u.HashID, "displayName": u.DisplayName, "tags": tags}

	return c.JSON(http.StatusCreated, data)
}

// CreateTag for getting user info by ID
func (h *UserHandler) CreateTag(c echo.Context) (err error) {
	var user model.User
	var tags []model.Tag
	body := map[string]interface{}{}

	if err = c.Bind(&body); err != nil {
		return
	}

	h.repo.CreateTag(&user, &tags, body)

	data := map[string]interface{}{"tags": tags}

	return c.JSON(http.StatusCreated, data)
}
