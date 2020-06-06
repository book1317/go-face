package model

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"errors"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var database_name = "facebook"
var collection_name = "Profile"

type Profilee struct {
	Client *mongo.Client
}

type Profile struct {
	ObjectID  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Image     string             `json:"image" bson:"image"`
}

func (p Profilee) GetProfiles(w http.ResponseWriter, r *http.Request) {
	Profile, err := getProfilesDB(p.Client)
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

func (p Profilee) GetProfileById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//w.WriteHeader(http.StatusOK)
	id, _ := vars["id"]
	Profile, err := getProfileByIdDB(p.Client, id)

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
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	//err := col.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)
	err := col.FindOne(ctx, bson.M{"_id": docID}).Decode(&result)

	// if err != nil {
	// 	fmt.Println("Error calling getProfileByIdDB:", err)
	// }
	return result, err
}

func (p Profilee) CreateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	var u Profile
	var a Account

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	k, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// handler error
	}

	if err := json.Unmarshal(k, &u); err != nil {
		// handle error
	}

	if err := json.Unmarshal(k, &a); err != nil {
		// handle error
	}

	// _ = json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("Profile===>", u)
	fmt.Println("Account===>", a)

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := p.Client.Database("facebook").Collection("profile")
	result, _ := col.InsertOne(ctx, u)
	json.NewEncoder(w).Encode(result)

	// oid, _ := result.InsertedID.(primitive.ObjectID)
	// fmt.Println("oid====>", oid)
	// a.ProfileID = oid

	// col2 := p.Client.Database("facebook").Collection("account")
	// result, _ = col2.InsertOne(ctx, a)
	json.NewEncoder(w).Encode(result)
}

func createProfileDB(client *mongo.Client, Profile Profile) error {
	return nil
}

func TestError() (string, error) {
	return "eiei", errors.New("this is an error")
}
