package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/manavp1996/mediumapi/web"
	"upper.io/db.v3/postgresql"
)

var settings = postgresql.ConnectionURL{
	Database: `postgres`,
	Host:     `db`,
	User:     `postgres`,
	Password: `password`,
}

func main() {

	sess, err := postgresql.Open(settings)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer sess.Close()
	collects, err := sess.Collections()
	if err != nil {
		log.Fatal("collections: ", err)
	}
	fmt.Printf("Collections: %v\n", collects)
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello", web.GetHello)
	router.Get("/users", web.GetUsers)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", router)

}
