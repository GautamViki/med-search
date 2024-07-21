package repository

import m "main.fo/model"

type UserRepo interface {
	GetUser()
	CreateUser() m.User
}
