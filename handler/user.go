package handler

import (
	"fmt"
	"medsearch/config"
	"medsearch/internal/dto"
	"medsearch/repository"
	"medsearch/repository/repo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type user struct {
	userRepo repository.UserRepo
}

func NewUserHandler() *user {
	return &user{
		userRepo: repo.NewUserRepo(),
	}
}

func (u *user) Get(c *gin.Context) {
	db := config.ConnectDB()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Invalid Id ", err.Error())
		return
	}
	user, err := u.userRepo.Get(id, db)
	if err != nil {
		fmt.Println("Got error while finding User ", err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (u *user) Create(c *gin.Context) {
	var req dto.UserRequest
	db := config.ConnectDB()
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := u.userRepo.Create(req, db)
	if err != nil {
		fmt.Println("ddddddddddddddddddddddddddddd", err.Error())
		return
	}
	// In a real application, you would save the user to a database
	c.JSON(http.StatusCreated, user)
}
