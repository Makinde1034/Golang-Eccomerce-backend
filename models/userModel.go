package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" validate:"email,required"`
	Password  string `json:"password" validate:"required"`
	Verified  bool
	ID        primitive.ObjectID `bson:"_id,omitempty"`
}