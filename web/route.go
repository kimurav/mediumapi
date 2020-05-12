package web

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/manavp1996/mediumapi/data"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user")

	w.Header().Set("Content-Type", "application/json")
	user := data.User{
		Id:    123,
		Name:  userName,
		Email: userName + "@mail.Sfjihsjdbasjks.com",
		Phone: 321321321,
	}

	json.NewEncoder(w).Encode(user)

}
