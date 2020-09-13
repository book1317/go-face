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

	profile, err := getProfileByAccountDB(db.Client, *account)
	if err != nil {
		fmt.Println("error ===> no document")
		return c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
		})
	}
	return c.JSON(http.StatusOK, profile)
}

func getProfileByAccountDB(client *mongo.Client, account Account) (*Profile, error) {
	var profile Profile
	var accountWithProfileID Account

	col := client.Database(db_facebook).Collection("account")
	err := col.FindOne(context.TODO(), bson.M{"username": account.Username}).Decode(&accountWithProfileID)
	if err != nil {
		return nil, err
		fmt.Println("error ===> no account")
	}

	fmt.Printf("accountWithProfileID ====> %+v", accountWithProfileID)
	col = client.Database(db_facebook).Collection("profile")
	// profileID, _ := primitive.ObjectIDFromHex(accountWithProfileID.ProfileID)
	// profileID, _ := primitive.ObjectIDFromHex("5f5a6a696f33b160afe72452")
	err = col.FindOne(context.TODO(), bson.M{"_id": accountWithProfileID.ProfileID}).Decode(&profile)
	if err != nil {
		return nil, err
		fmt.Println("error ===> no profile")
	}
	return &profile, err
}
