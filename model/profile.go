package model

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var database_name = "facebook"
var collection_name = "profile"

type Profile struct {
	Client *mongo.Client
}

type profile struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Image     string `json:"image"`
}

func (p Profile) GetProfiles(w http.ResponseWriter, r *http.Request) {
	profile := getProfiles_DB(p.Client)
	//w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	json.NewEncoder(w).Encode(profile)
	fmt.Println("Get Profiles")
}

func (p Profile) GetProfileByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//w.WriteHeader(http.StatusOK)
	id, _ := strconv.Atoi(vars["id"])
	profile := getProfileByID_DB(p.Client, id)

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	json.NewEncoder(w).Encode(profile)
	fmt.Println("GetProfileByID :", id)
}

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p Profile) GetProfileByAccount(w http.ResponseWriter, r *http.Request) {

	//profile := profiles_mock()[0]

	var u Account
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		//http.Error(w, err.Error(), 400)

		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Println(u)

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	// json.NewEncoder(w).Encode(r)
	fmt.Fprintf(w, "Login")
	//fmt.Println("GetProfileByID :", id)
}

func profiles_mock() []profile {
	profiles := []profile{
		{
			ID:        1,
			Firstname: "Chaiyarin",
			Lastname:  "Niamsuwan",
			Image:     "/static/media/profile1.6faebd7c.png",
		},
		{
			ID:        2,
			Firstname: "BooKy",
			Lastname:  "Eiei",
			Image:     "/static/media/profile1.6faebd7c.png",
		},
	}
	return profiles
}

func getProfiles_DB(client *mongo.Client) []profile {
	var result []profile
	col := client.Database(database_name).Collection(collection_name)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(ctx)
	} else {
		// iterate over docs using Next()
		for cursor.Next(ctx) {
			// Declare a result BSON object
			var res bson.M
			err := cursor.Decode(&res) // If there is a cursor.Decode error
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
			} else {
				var profile profile
				bsonBytes, _ := bson.Marshal(res)
				bson.Unmarshal(bsonBytes, &profile)
				result = append(result, profile)

				// fmt.Println(result)
			}
		}
	}
	return result
}

func getProfileByID_DB(client *mongo.Client, id int) profile {
	var result profile
	col := client.Database("facebook").Collection("profile")
	//docID, _ := primitive.ObjectIDFromHex("5e93293d77d2db0f4249b3bb")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	//err := col.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)
	err := col.FindOne(ctx, bson.M{"id": id}).Decode(&result)

	if err != nil {
		fmt.Println("Error calling getProfileByID_DB:", err)
	} else {
		// fmt.Println("FindOne() result:", result)
	}
	return result
}
