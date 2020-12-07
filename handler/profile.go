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
	profile, err := model.GetProfileByIdDB(db.Client, profileId)
	if err != nil {
		fmt.Println("profileId ====> ", c.Param("id"))
		fmt.Println("error ====> GetProfileByIdDB")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, gin.H{"data": profile})
}

func (db Database) UpdateProfileImageById(c echo.Context) error {
	fmt.Println("UpdateProfileImageById")

	profileId := c.Param("id")
	profileImage := model.Image{}
	// fmt.Printf("c.Request.Body ===> %+v", c.Request)

	fmt.Println("profileIdd", profileId)
	if err := c.Bind(&profileImage); err != nil {
		fmt.Println("error ===> no payload")
		return c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
	}

	profileIdd, err := model.UpdateProfileImageDB(db.Client, profileId, profileImage)
	if err != nil {
		fmt.Println("error ====> GetProfileByIdDB")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, gin.H{"data": profileIdd})
}
