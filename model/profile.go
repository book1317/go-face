package model

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"reflect"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var database_name = db_facebook
var collection_name = "Profile"

type Profile struct {
	ObjectID  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Image     string             `json:"image" bson:"image"`
}

func (db Database) GetProfiles(w http.ResponseWriter, r *http.Request) {
	Profile, err := getProfilesDB(db.Client)
	//w.WriteHeader(http.StatusOK)
	if err != nil {
		fmt.Fprintf(w, "Not found")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		json.NewEncoder(w).Encode(Profile)
		fmt.Println("Get Profiles")
	}
}

func (db Database) GetProfileById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]
	Profile, err := getProfileByIdDB(db.Client, id)
	if err != nil {
		fmt.Fprintf(w, "Not found")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		json.NewEncoder(w).Encode(Profile)
		fmt.Println("GetProfileById :", id)
	}
}

func getProfilesDB(client *mongo.Client) ([]Profile, error) {
	var result []Profile
	col := client.Database(database_name).Collection(collection_name)
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

func getProfileByIdDB(client *mongo.Client, id string) (Profile, error) {
	var result Profile
	col := client.Database(db_facebook).Collection("profile")
	docID, _ := primitive.ObjectIDFromHex(id)
	err := col.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)
	return result, err
}

func (db Database) insertProfileDB(profile Profile) (primitive.ObjectID, error) {
	col := db.Client.Database(db_facebook).Collection("profile")
	result, err := col.InsertOne(context.TODO(), profile)
	fmt.Println(reflect.TypeOf(result))
	profileID, _ := result.InsertedID.(primitive.ObjectID)
	return profileID, err
}
