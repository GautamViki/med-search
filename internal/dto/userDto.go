package dto

import httpresponse "medsearch/helper/httpResponse"

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

type UserResponse struct {
	httpresponse.Response
	User User `json:"user"`
}
