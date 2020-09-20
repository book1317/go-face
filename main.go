package main

import (
	"go-face/database"
	"go-face/handler"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Client *mongo.Client
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut, http.MethodPatch},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Disposition"},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := database.Database{}
	client := database.Connect()
	db := handler.Database{client}

	e.POST("/register", db.Register)
	e.POST("/login", db.Login)
	e.POST("/posts", db.CreatePost)
	e.GET("/posts", db.GetPosts)
	e.GET("/profile/:id", db.GetProfileById)
	e.PATCH("/profile/image/:id", db.UpdateProfileImageById)
	e.PATCH("/comment/:id", db.InserComment)
	e.Logger.Fatal(e.Start(":" + "8080"))
}
