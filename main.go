package main

import (
	"go-face/database"
	"go-face/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut, http.MethodPatch},
		AllowHeaders: []string{"*"},
		//[]string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Disposition"},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := database.Database{}
	client := database.Connect()
	db := model.Database{client}

	e.POST("/register", db.Register)
	e.POST("/login", db.Login)
	e.POST("/post", db.CreatePost)
	e.GET("/post", db.GetPosts)
	e.PATCH("/comment/:id", db.InserComment)
	e.Logger.Fatal(e.Start(":" + "8080"))
}
