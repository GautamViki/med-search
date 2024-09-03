package config

import (
	"medsearch/model"

	"gorm.io/gorm"
)

func UpdateDB(db *gorm.DB) {
	db.AutoMigrate(model.User{})
}
