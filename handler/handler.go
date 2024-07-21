package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"main.fo/repository"
	v1 "main.fo/repository/v1"
)

type UserRepo struct {
	userRepo repository.UserRepo
}

func NewUserHandler() *UserRepo {
	return &UserRepo{
		userRepo: v1.NewUserRepo(),
	}
}

// HomeHandler handles the root route
func (u *UserRepo) Create(w http.ResponseWriter, r *http.Request) {
	user := u.userRepo.CreateUser()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// HelloHandler handles the /hello route
func Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// ListUsersHandler handles the /users route
func GetById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

// GetUserHandler handles the /users/{userID} route
func Update(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Write([]byte("User ID: " + userID))
}
