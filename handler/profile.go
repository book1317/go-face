package handler

import (
	"encoding/json"
	"fmt"
	"go-face/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

func (db Database) GetProfiles(w http.ResponseWriter, r *http.Request) {
	Profile, err := model.GetProfilesDB(db.Client)
	if err != nil {
		fmt.Fprintf(w, "Not found")
	} else {
		json.NewEncoder(w).Encode(Profile)
	}
}

func (db Database) GetProfileById(c echo.Context) error {
	profileId := c.Param("id")
	fmt.Println("profileId ===>", profileId)

	profile, err := model.GetProfileByIdDB(db.Client, profileId)
	if err != nil {
		fmt.Println("error ====> GetProfileByIdDB")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, gin.H{"data": profile})
}
