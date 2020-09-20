package model

import (
	"context"
	"fmt"
	"log"

	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Profile struct {
	ObjectID  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Image     string             `json:"image" bson:"image"`
}

type Image struct {
	Image string `json:"image" bson:"image"`
}

func GetProfilesDB(client *mongo.Client) ([]Profile, error) {
	var result []Profile
	col := client.Database(db_facebook).Collection(co_profile)
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(context.TODO())
	} else {
		for cursor.Next(context.TODO()) {
			var res bson.M
			err := cursor.Decode(&res)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
			} else {
				var Profile Profile
				bsonBytes, _ := bson.Marshal(res)
				bson.Unmarshal(bsonBytes, &Profile)
				result = append(result, Profile)
			}
		}
	}
	return result, err
}

func GetProfileByIdDB(client *mongo.Client, id string) (*Profile, error) {
	var profile Profile
	col := client.Database(db_facebook).Collection(co_profile)
	docID, _ := primitive.ObjectIDFromHex(id)
	err := col.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&profile)
	if err != nil {
		fmt.Println("error cannot find document")
		return nil, err
	}
	return &profile, err
}

func InsertProfileDB(client *mongo.Client, profile Profile) (primitive.ObjectID, error) {
	col := client.Database(db_facebook).Collection(co_profile)
	result, err := col.InsertOne(context.TODO(), profile)
	fmt.Println(reflect.TypeOf(result))
	profileID, _ := result.InsertedID.(primitive.ObjectID)
	return profileID, err
}

func UpdateProfileImageDB(client *mongo.Client, profileId string, image Image) (*primitive.ObjectID, error) {
	col := client.Database(db_facebook).Collection(co_profile)
	id, _ := primitive.ObjectIDFromHex(profileId)
	result, err := col.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"image", image.Image}}},
		},
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	profileID, _ := result.UpsertedID.(primitive.ObjectID)
	return &profileID, err
}
