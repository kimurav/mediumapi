package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/manavp1996/mediumapi/data"
	"github.com/manavp1996/mediumapi/web"
)

func main() {

	data.SetupDB()
	defer data.DB.Close()
	collects, err := data.DB.Collections()
	if err != nil {
		log.Fatal("collections: ", err)
	}
	fmt.Printf("Collections: %v\n", collects)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)

	router.Route("/api/user", func(r chi.Router) {
		r.Post("/", web.CreateUser)
	})

	http.ListenAndServe(":8080", router)

}
