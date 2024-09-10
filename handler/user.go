package handler

import (
	"encoding/json"
	"fmt"
	"medsearch/config"
	"medsearch/helper"
	httpresponse "medsearch/helper/httpResponse"
	"medsearch/internal/dto"
	"medsearch/model"
	"medsearch/repository"
	"medsearch/repository/repo"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type user struct {
	userRepo repository.UserRepo
}

func NewUserHandler() *user {
	return &user{
		userRepo: repo.NewUserRepo(),
	}
}

func (u *user) Get(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid Id ", err.Error())
		return
	}
	user, err := u.userRepo.Get(id, db)
	if err != nil {
		fmt.Println("Got error while finding User ", err.Error())
		return
	}
	userDto := createUserDto(user, "0", "User fetched successfully.")
	helper.RespondWithJSON(w, userDto, http.StatusOK)
	return
}

func (u *user) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.UserRequest
	db := config.ConnectDB()
	// Decode the incoming JSON request body into the user struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := httpresponse.PrepareResponse("1000", "Error occurred while decoding body.")
		helper.RespondWithError(w, http.StatusBadRequest, res)
		return
	}
	user, err := u.userRepo.Create(req, db)
	if err != nil {
		res := httpresponse.PrepareResponse("2002", "Error occurred while creating user.")
		helper.RespondWithError(w, http.StatusInternalServerError, res)
		return
	}
	userDto := createUserDto(user, "0", "User created successfully.")
	helper.RespondWithJSON(w, userDto, http.StatusCreated)
	return
}

func createUserDto(user model.User, code, message string) dto.UserResponse {
	userDto := dto.User{
		EntityId: user.EntityId,
		Name:     user.Name,
		Address:  user.Address,
		Email:    user.Email,
		DOb:      user.Dob,
	}
	return dto.UserResponse{
		Response: httpresponse.Response{
			Code:    code,
			Message: message,
		},
		User: userDto,
	}
}
