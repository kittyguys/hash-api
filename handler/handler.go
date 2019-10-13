package handler

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/kittyguys/hash/api/config"
	"github.com/labstack/echo"
)

var router *mux.Router

type (
	// Handler Handler
	Handler struct {
		DB *gorm.DB
	}
)

const (
	// Key This should be imported from somewhere else
	Key = "secret"
)

// InitializeRouter Init Router
func InitializeRouter(db *gorm.DB, e *echo.Echo) *echo.Echo {
	config := config.New()
	user := NewUserHandler(db)
	tag := NewTagHandler(db)
	auth := NewOAuthHandler(config, db)

	// Auth
	e.POST("/signup", user.SignUp)
	e.POST("/login", user.Login)
	// Social login
	e.GET("/auth/twitter", auth.TwitterLogin())
	e.GET("/auth/twitter/callback", auth.TwitterCallback())
	// User
	e.GET("/users/:id", user.GetUser)
	e.POST("/users/:id/tags", user.CreateTag)
	// Tag
	e.GET("/tags/:name/users", tag.GetUsers)
	return e
}
