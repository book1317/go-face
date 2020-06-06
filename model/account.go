package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ProfileID primitive.ObjectID `json:"profileID" bson:"profile_id"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
}

// func GetProfileByAccountDB(client *mongo.Client, acc Account) (Profile, error) {
// 	var result Profile
// 	col := client.Database("facebook").Collection("profile")
// 	docID, _ := primitive.ObjectIDFromHex(id)
// 	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
// 	err := col.FindOne(ctx, bson.M{"_id": docID}).Decode(&result)

// 	return result, err
// }

// func profiles_mock() []profile {
// 	profiles := []profile{
// 		{
// 			ID:        "1",
// 			Firstname: "Chaiyarin",
// 			Lastname:  "Niamsuwan",
// 			Image:     "/static/media/profile1.6faebd7c.png",
// 		},
// 		{
// 			ID:        "2",
// 			Firstname: "BooKy",
// 			Lastname:  "Eiei",
// 			Image:     "/static/media/profile1.6faebd7c.png",
// 		},
// 	}
// 	return profiles
// }
