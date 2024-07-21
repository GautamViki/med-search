package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	handlers "main.fo/handler"
	"main.fo/model"
)

func main() {
	router := chi.NewRouter()
	fmt.Println("server started")
	dsn := "root:root@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	fmt.Println("db connected")
	db.AutoMigrate(&model.User{})
	// Use some middlewares
	router.Use(middleware.Logger)

	hand := handlers.NewUserHandler()

	router.Route("/users", func(r chi.Router) {
		r.Post("/", hand.Create)
		r.Get("/hello", handlers.Get)
		r.Get("/", handlers.GetById)
		r.Put("/{userID}", handlers.Update)
	})
	http.ListenAndServe(":3009", router)
}
