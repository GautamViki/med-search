package repo

import (
	"medsearch/internal/dto"
	"medsearch/model"
	"medsearch/repository"

	"gorm.io/gorm"
)

type mysqlUserRepo struct {
	userRepo repository.UserRepo
}

func NewUserRepo() repository.UserRepo {
	return &mysqlUserRepo{}
}

func (u *mysqlUserRepo) Get(id int, db *gorm.DB) (model.User, error) {
	var user model.User
	if err := db.Where("entity_id = ?", id).Find(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *mysqlUserRepo) Create(req dto.UserRequest, db *gorm.DB) (model.User, error) {
	user := model.User{
		Name:    req.Name,
		Dob:     req.DOb,
		Address: req.Address,
		Email:   req.Email,
	}
	if err := db.Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
