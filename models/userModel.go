package models

type User struct {
	Firstname string `json:"firstname" validate:"required,min=2,max=30"`
	Lastname string `json:"lastname" validate:"required,min=2,max=30"`
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
	verified bool
}