package handler

import (
	"go-face/app"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type route struct {
	Desc           string
	Group          string
	Path           string
	HttpMethod     string
	HandlerFunc    echo.HandlerFunc
	MiddlewareFunc []echo.MiddlewareFunc
}

func NewRouter(e *echo.Echo, c *app.Config) error {
	// e.POST("/login", db.Login)
	// e.POST("/posts", db.CreatePost)
	// e.GET("/posts", db.GetPosts)
	// e.GET("/profile/:id", db.GetProfileById)
	// e.PATCH("/profile/image/:id", db.UpdateProfileImageById)
	// e.PATCH("/comment/:id", db.InserComment)

	db := Database{c.MongoDB.Client}
	routes := []route{
		{
			Desc:        "Register",
			Path:        "/register",
			HttpMethod:  http.MethodPost,
			HandlerFunc: db.Register,
		},
		{
			Desc:        "Login",
			Path:        "/login",
			HttpMethod:  http.MethodPost,
			HandlerFunc: db.Login,
		},
		{
			Desc:        "Create Post",
			Path:        "/posts",
			HttpMethod:  http.MethodPost,
			HandlerFunc: db.CreatePost,
		},
		{
			Desc:        "Get Posts",
			Path:        "/posts",
			HttpMethod:  http.MethodGet,
			HandlerFunc: db.GetPosts,
		},
		{
			Desc:        "Get Profile By Id",
			Group:       "profile",
			Path:        "/:id",
			HttpMethod:  http.MethodGet,
			HandlerFunc: db.GetProfileById,
		},
		{
			Desc:        "Inser Comment",
			Group:       "comment",
			Path:        "/:id",
			HttpMethod:  http.MethodPatch,
			HandlerFunc: db.InserComment,
		},
		{
			Desc:        "Update Profile Image By Id",
			Group:       "profile",
			Path:        "/image/:id",
			HttpMethod:  http.MethodPatch,
			HandlerFunc: db.UpdateProfileImageById,
		},
	}

	// middleware
	// e.Use(middleware.BodyDumpWithConfig(bodyDumpConfig()))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut, http.MethodPatch},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Disposition"},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	for _, r := range routes {
		e.Group(r.Group).Add(r.HttpMethod, r.Path, r.HandlerFunc, r.MiddlewareFunc...)
	}

	return nil
}
