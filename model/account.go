package model

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Account struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func (p Profilee) CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var u Profile

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	fmt.Println("CreateProfile===>", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&u)

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := p.Client.Database("facebook").Collection("profile")
	result, insertErr := col.InsertOne(ctx, u)
	if insertErr != nil {
		fmt.Println(insertErr)
	} else {
		fmt.Println(result)
	}

	json.NewEncoder(w).Encode(result)
}

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
