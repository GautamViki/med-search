package model

type User struct {
	Name    string `json:"name" gorm:"primaryKey"`
	Address string
}
