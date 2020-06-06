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

	// e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut, http.MethodPatch},
	// 	AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
	// 	AllowCredentials: true,
	// 	ExposeHeaders:    []string{"Content-Disposition"},
	// }))
	// e.Use(middleware.Logger())

	// Enable recover APIs panic
	// TODO: re-enable when push
	// e.Use(middleware.Recover())

	profile := model.Profilee{client}

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/profiles", profile.GetProfiles).Methods("GET", "OPTIONS")
	r.HandleFunc("/profiles/{id}", profile.GetProfileById).Methods("GET", "OPTIONS")
	//	r.HandleFunc("/login", profile.GetProfileByAccount).Methods("POST")
	r.HandleFunc("/register", profile.CreateProfile).Methods("POST", "OPTIONS")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// e.Logger.Fatal(e.Start(":" + "8080"))
}
