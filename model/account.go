package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	ProfileID primitive.ObjectID `json:"profileID" bson:"profile_id"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
}

func (db Database) InserAccountDB(account Account) (*mongo.InsertOneResult, error) {
	col := db.Client.Database("facebook").Collection("account")
	result, err := col.InsertOne(context.TODO(), account)
	return result, err
}
