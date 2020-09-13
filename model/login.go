package model

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db Database) Login(c echo.Context) error {
	fmt.Println("==== Longin ====")
	account := new(Account)

	if err := c.Bind(account); err != nil {
		fmt.Println("error ===> no payload")
		return c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
	}

	profile, err := getProfileByUsernameDB(db.Client, *account)
	if err != nil {
		fmt.Println("error ===> no document")
		return c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
		})
	}
	return c.JSON(http.StatusOK, profile)
}

func getProfileByUsernameDB(client *mongo.Client, account Account) (*Profile, error) {
	var profile Profile
	var accountWithProfileID Account

	col := client.Database(db_facebook).Collection(co_account)
	err := col.FindOne(context.TODO(), bson.M{"username": account.Username}).Decode(&accountWithProfileID)
	if err != nil {
		fmt.Println("error ===> no account")
		return nil, err
	}

	if account.Password != accountWithProfileID.Password {
		fmt.Println("error ===> wrong password")
		return nil, err
	}

	col = client.Database(db_facebook).Collection(co_profile)
	err = col.FindOne(context.TODO(), bson.M{"_id": accountWithProfileID.ProfileID}).Decode(&profile)
	if err != nil {
		fmt.Println("error ===> no profile")
		return nil, err
	}
	return &profile, err
}
