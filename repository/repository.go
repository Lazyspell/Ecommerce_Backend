package repository

import "github.com/lazyspell/Ecommerce_Backend/models"

type DatabaseRepo interface {
	AllCategories() ([]models.Categories, error)
	CategoryById(id int) (models.Categories, error)

	NewUserDB(user models.Users) (string, error)
	// GetPassword(id int) (models.Users, error)
	AllUsers() ([]models.DisplayUser, error)
	UserById(id int) (models.DisplayUser, error)
	DeleteUserDB(id int) (string, error)

	NewGoogleUserDB(googleUser models.GoogleObject) (string, error)
	GoogleAuthenticate(email string) (models.GoogleObject, error)

	Authenticate(email string) (models.Users, error)
}
