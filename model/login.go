package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProfileByUsernameDB(client *mongo.Client, account Account) (*string, error) {
	// var profile Profile
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

	// col = client.Database(db_facebook).Collection(co_profile)
	// err = col.FindOne(context.TODO(), bson.M{"_id": accountWithProfileID.ProfileID}).Decode(&profile)
	// if err != nil {
	// 	fmt.Println("error ===> no profile")
	// 	return nil, err
	// }
	profileId := accountWithProfileID.ProfileID.Hex()
	return &profileId, err
}
