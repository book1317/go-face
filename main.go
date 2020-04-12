package main

import (
	"fmt"
	"go-face/database"
	"go-face/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "404", r.URL.Path[1:])
}

func main() {
	database := database.Database{}
	client := database.Connect()

	//collection := client.Database("facebook").Collection("profile")
	// data := models.PROFILES_MOCK()[0]
	// insertResult, err := collection.InsertOne(context.TODO(), data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

	profile := model.Profile{client}

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/get_profiles", profile.GetProfiles).Methods("GET")
	r.HandleFunc("/get_profiles/{id}", profile.GetProfileByID).Methods("GET")
	r.HandleFunc("/login", profile.GetProfileByAccount).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
