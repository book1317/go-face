package main

import (
	"flag"
	"fmt"
	"go-face/app"
	"go-face/handler"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Client *mongo.Client
}

func main() {
	state := flag.String("state", "dev", "program environment")
	flag.Parse()

	c := app.NewConfig(*state)
	if err := c.Init(); err != nil {
		fmt.Println("config init error", err)
		return
	}

	e := echo.New()
	if err := handler.NewRouter(e, c); err != nil {
		fmt.Println("new router error", err)
		return
	}

	e.Start(":" + "8080")
}
