package model

import (
	"net/http"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p Profile) GetProfileByAccount(w http.ResponseWriter, r *http.Request) {

	// profile := profiles_mock()[0]

	// fmt.Println("Login")
	// var u Account
	// if r.Body == nil {
	// 	http.Error(w, "Please send a request body", 400)
	// 	return
	// }
	// err := json.NewDecoder(r.Body).Decode(&u)
	// if err != nil {
	// 	fmt.Fprintf(w, "Not found")
	// }

	// w.Header().Add("Content-Type", "application/json")
	// w.Header().Add("Access-Control-Allow-Origin", "*")
	// w.Header().Add("Access-Control-Allow-Headers", "*")

	// json.NewEncoder(w).Encode(profile)

	// json.NewEncoder(w).Encode(r)
	//fmt.Fprintf(w, "Login")
	//fmt.Println("GetProfileById :", id)
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
