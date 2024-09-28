package main

import (
	"fmt"
	"medsearch/config"
	"medsearch/handler"
	"net/http"
	"os"

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
	fmt.Println("hiiii")
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Println("Server started at port", port)
	http.ListenAndServe(port, r)
}
