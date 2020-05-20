package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/manavp1996/mediumapi/data"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user data.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
		return
	}
	log.Printf("%+v\n", user)
	if _, err := data.DB.InsertInto("users").Values(user).Exec(); err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
		return
	}
	w.WriteHeader(200)
}
