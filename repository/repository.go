package repository

import "github.com/lazyspell/Ecommerce_Backend/models"

type DatabaseRepo interface {
	AllCategories() ([]models.Categories, error)
	CategoryById(id int) (models.Categories, error)
	NewUserDB(user models.Users) (string, error)
	GetPassword(id int) (models.Users, error)
	AllUsers() ([]models.Users, error)
	UserById(id int) (models.Users, error)
	DeleteUserDB(id int) (string, error)
}
