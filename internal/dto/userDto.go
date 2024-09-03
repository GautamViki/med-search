package dto

type UserRequest struct {
	Name    string `json:"name"`
	DOb     string `json:"dob"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

type User struct {
	EntityId int    `json:"entityId"`
	Name     string `json:"name"`
	DOb      string `json:"dob"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}
