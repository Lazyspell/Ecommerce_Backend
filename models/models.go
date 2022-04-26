package models

type Categories struct {
	Id       int    `json:"id"`
	Title    string `json:"category"`
	ImageUrl string `json:"image_url"`
}
