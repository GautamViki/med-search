package model

type User struct {
	EntityId int    `gorm:"primaryKey;type:INT(10)" json:"entity_id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Dob      string `gorm:"size:100" json:"dob"`
	Address  string `gorm:"size:100" json:"address"`
	Email    string `gorm:"size:100" json:"email"`
}
