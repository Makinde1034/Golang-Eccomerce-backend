package models

import (
	"time"
)

type item struct {
	Name        string             `json:"firstname"`
	Price      	string             `json:"price"`
	Description string             `json:"description"`
	
}

type Store struct {
	Name        string             `json:"name"`
	Created     time.Time          `json:"created"`
	Description string             `json:"description"`
	Category    string             `json:"category"`
	Owner       string             `json:"owner"`
	Orders      string             `json:"orders"`
	Products    []item             `json:"products"`
	Image       string             `json:"image"`
	PhoneNumber string             `json:"phoneNumber"`
	Whatsapp    string             `json:"whatsapp"`
}