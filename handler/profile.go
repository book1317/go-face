package handler

import (
	"encoding/json"
	"fmt"
	"go-face/model"
	"net/http"

	"github.com/gorilla/mux"
)

func (db Database) GetProfiles(w http.ResponseWriter, r *http.Request) {
	Profile, err := model.GetProfilesDB(db.Client)
	if err != nil {
		fmt.Fprintf(w, "Not found")
	} else {
		json.NewEncoder(w).Encode(Profile)
	}
}

func (db Database) GetProfileById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]
	Profile, err := model.GetProfileByIdDB(db.Client, id)
	if err != nil {
		fmt.Fprintf(w, "Not found")
	} else {
		json.NewEncoder(w).Encode(Profile)
	}
}
