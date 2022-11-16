package models

type Product struct {
	Name string `json:"name"`
	Price string `json:"price"`
	Category string `json:"category"`
	Color string `json:"color"`
	Description string `json:"description"`
	Amount string `json:"amount"`
	Images []string	`json:"images"`
	// Owner string `json:"Owner"`
}