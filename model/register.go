package model

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Register struct {
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
}

func (db Database) Register(c echo.Context) error {
	register := new(Register)
	account := Account{}
	profile := Profile{}

	if err := c.Bind(register); err != nil {
		return err
	}

	profile.Firstname = register.Firstname
	profile.Lastname = register.Lastname
	profileID, err := db.insertProfileDB(profile)
	if err != nil {
		return err
	}

	account.Username = register.Username
	account.Password = register.Password
	account.ProfileID = profileID
	result, err := db.InserAccountDB(account)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}
