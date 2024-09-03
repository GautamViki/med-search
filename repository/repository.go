package repository

import (
	"medsearch/internal/dto"
	"medsearch/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(req dto.UserRequest, db *gorm.DB) (model.User, error)
	Get(id int, db *gorm.DB) (model.User, error)
}
