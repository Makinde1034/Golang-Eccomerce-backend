package models

type User struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
	Verified bool
}