package handler

import (
	"go-face/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (db Database) Register(c echo.Context) error {
	register := new(model.Register)
	account := model.Account{}
	profile := model.Profile{}

	if err := c.Bind(register); err != nil {
		return err
	}

	profile = register.Profile
	profileID, err := model.InsertProfileDB(db.Client, profile)
	if err != nil {
		return err
	}

	account = register.Account
	account.ProfileID = profileID
	result, err := model.InserAccountDB(db.Client, account)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}
