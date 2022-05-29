package models

import "time"

type Categories struct {
	Id       int    `json:"id"`
	Title    string `json:"category"`
	ImageUrl string `json:"image_url"`
}

type Products struct {
	Id          int       `json:"id"`
	ProductName string    `json:"hat_name"`
	ImageUrl    string    `json:"image_url"`
	Price       int       `json:"price"`
	Product     string    `json:"product"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
