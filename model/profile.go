package model

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"errors"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var database_name = "facebook"
var collection_name = "profile"

type Profile struct {
	Client *mongo.Client
}

type profile struct {
	ObjectID  primitive.ObjectID `json:"id" bson:"_id"`
	Firstname string             `json:"firstname"`
	Lastname  string             `json:"lastname"`
	Image     string             `json:"image"`
}

func (p Profile) GetProfiles(w http.ResponseWriter, r *http.Request) {
	profile, err := getProfilesDB(p.Client)
	//w.WriteHeader(http.StatusOK)
	if err != nil {
		fmt.Fprintf(w, "Not found")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		json.NewEncoder(w).Encode(profile)
		fmt.Println("Get Profiles")
	}
}

func (p Profile) GetProfileById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//w.WriteHeader(http.StatusOK)
	id, _ := vars["id"]
	profile, err := GetProfileByIdDB(p.Client, id)

	if err != nil {
		fmt.Fprintf(w, "Not found")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		json.NewEncoder(w).Encode(profile)
		fmt.Println("GetProfileById :", id)
	}

}

func getProfilesDB(client *mongo.Client) ([]profile, error) {
	var result []profile
	col := client.Database(database_name).Collection(collection_name)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(ctx)
	} else {
		for cursor.Next(ctx) {
			var res bson.M
			err := cursor.Decode(&res) // If there is a cursor.Decode error
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
			} else {
				var profile profile
				bsonBytes, _ := bson.Marshal(res)
				bson.Unmarshal(bsonBytes, &profile)
				result = append(result, profile)
			}
		}
	}
	return result, err
}

func GetProfileByIdDB(client *mongo.Client, id string) (profile, error) {
	var result profile
	col := client.Database("facebook").Collection("profile")
	docID, _ := primitive.ObjectIDFromHex(id)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	//err := col.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)
	err := col.FindOne(ctx, bson.M{"_id": docID}).Decode(&result)

	// if err != nil {
	// 	fmt.Println("Error calling GetProfileByIdDB:", err)
	// }
	return result, err
}

func TestError() (string, error) {
	return "eiei", errors.New("this is an error")
}
