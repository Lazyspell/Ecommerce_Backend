package models

import "time"

type Categories struct {
	Id       int    `json:"id"`
	Title    string `json:"category"`
	ImageUrl string `json:"image_url"`
}

type Products struct {
	Id          int       `json:"id"`
	ProductName string    `json:"product_name"`
	ImageUrl    string    `json:"image_url"`
	Price       int       `json:"price"`
	Product     string    `json:"product"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Items struct {
	Hats     []Products `json:"hats"`
	Sneakers []Products `json:"sneakers"`
	Jackets  []Products `json:"jackets"`
	Womens   []Products `json:"womens"`
	Mens     []Products `json:"mens"`
}

type ProductsObject struct {
	objects []Items
}
