package v1

import (
	"fmt"

	m "main.fo/model"
	"main.fo/repository"
)

type mysqlUserRepo struct {
}

func NewUserRepo() repository.UserRepo {
	return &mysqlUserRepo{}
}

func (u *mysqlUserRepo) CreateUser() m.User {
	
	return m.User{Name: "vikas", Address: "UP"}
}

func (u *mysqlUserRepo) GetUser() {
	fmt.Println("GetUser repo function")
}
