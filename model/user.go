package model

type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}