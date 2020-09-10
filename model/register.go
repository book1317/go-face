package model

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Register struct {
	Profile Profile `json:"profile" bson:"profile"`
	Account Account `json:"account" bson:"account"`
}

func (db Database) Register(c echo.Context) error {
	register := new(Register)
	account := Account{}
	profile := Profile{}

	if err := c.Bind(register); err != nil {
		return err
	}

	profile = register.Profile
	profileID, err := db.insertProfileDB(profile)
	if err != nil {
		return err
	}

	account = register.Account
	account.ProfileID = profileID
	result, err := db.InserAccountDB(account)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}
