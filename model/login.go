package model

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db Database) Login(c echo.Context) error {
	fmt.Println("==== Longin ====")
	account := new(Account)

	if err := c.Bind(account); err != nil {
		fmt.Println("errrrrrrrrrrrrrrrrrr", err)
		return err
	}
	fmt.Printf("account ====> %+v", account)

	profile, err := getProfileByAccountDB(db.Client, *account)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusCreated, profile)
}

func getProfileByAccountDB(client *mongo.Client, account Account) (Profile, error) {
	var profile Profile
	var accountWithProfileID Account

	fmt.Printf("account ====> %+v", account)
	col := client.Database("facebook").Collection("account")
	err := col.FindOne(context.TODO(), bson.M{"$match": bson.M{"username": account.Username, "password": account.Password}}).Decode(&accountWithProfileID)
	if err != nil {
		// return nil, err
		fmt.Println("find account error")
	}

	col = client.Database("facebook").Collection("profile")
	// profileID, _ := primitive.ObjectIDFromHex(accountWithProfileID.ProfileID)
	err = col.FindOne(context.TODO(), bson.M{"_id": accountWithProfileID.ProfileID}).Decode(&profile)
	if err != nil {
		// return nil, err
		fmt.Println("find profile error")
	}
	return profile, err
}
