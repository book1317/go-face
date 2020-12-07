package handler

import (
	"fmt"
	"go-face/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

func (db Database) Login(c echo.Context) error {
	fmt.Println("==== Longin ====")
	account := new(model.Account)

	if err := c.Bind(account); err != nil {
		fmt.Println("error ===> no payload")
		return c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
	}

	profileId, err := model.GetProfileByUsernameDB(db.Client, *account)
	if err != nil {
		fmt.Println("error ===> no document")
		return c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
		})
	}
	return c.JSON(http.StatusOK, profileId)
}
