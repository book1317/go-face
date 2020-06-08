package model

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"errors"

	"reflect"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var database_name = "facebook"
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
	//w.WriteHeader(http.StatusOK)
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
			err := cursor.Decode(&res) // If there is a cursor.Decode error
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
	col := client.Database("facebook").Collection("profile")
	docID, _ := primitive.ObjectIDFromHex(id)
	// ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	//err := col.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)
	err := col.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)

	// if err != nil {
	// 	fmt.Println("Error calling getProfileByIdDB:", err)
	// }
	return result, err
}

func (db Database) CreateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	var profile Profile
	var account Account

	if r.Body == http.NoBody {
		//http.Error(w, "Please send a request body", 400)
		return
	}

	k, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// handler error
	}
	if err := json.Unmarshal(k, &profile); err != nil {
		// handle error
	}
	if err := json.Unmarshal(k, &account); err != nil {
		// handle error
	}

	// _ = json.NewDecoder(r.Body).Decode(&u)
	profileID, err := db.insertProfileDB(profile)
	json.NewEncoder(w).Encode(profileID)
	account.ProfileID = profileID

	//result, err := db.InserAccountDB(account)

	// db.CreateAccount()

	// 	col2 := db.Client.Database("facebook").Collection("account")
	// 	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	// 	result, _ := col2.InsertOne(ctx, account)
}

func (db Database) insertProfileDB(profile Profile) (primitive.ObjectID, error) {
	col := db.Client.Database("facebook").Collection("profile")
	result, err := col.InsertOne(context.TODO(), profile)
	fmt.Println(reflect.TypeOf(result))
	profileID, _ := result.InsertedID.(primitive.ObjectID)
	return profileID, err
}

func TestError() (string, error) {
	return "eiei", errors.New("this is an error")
}
