package main

import (
	"fmt"
	"medsearch/config"
	"medsearch/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	db := config.ConnectDB()
	config.UpdateDB(db)
	r := chi.NewRouter()
	user := handler.NewUserHandler()
	r.Group(func(router chi.Router) {
		router.Get("/users/{id}", user.Get)
		router.Post("/users", user.Create)
	})
	fmt.Println("Server started at port:3007")
	http.ListenAndServe(":3007", r)
}
