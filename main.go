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
	result, error := model.TestError()
	fmt.Println("result=====>", result)
	fmt.Println("error=====>", error)

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
	r.HandleFunc("/profiles", profile.GetProfiles).Methods("GET")
	r.HandleFunc("/profiles/{id}", profile.GetProfileById).Methods("GET")
	r.HandleFunc("/login", profile.GetProfileByAccount).Methods("POST")
	r.HandleFunc("/register", profile.CreateProfile).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
