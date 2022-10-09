package response

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterResponse struct {
	Firstname string
	Lastname  string
	Email     string
	Token     string
	ID        primitive.ObjectID `bson:"_id"`
}