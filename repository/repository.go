package repository

import "github.com/lazyspell/Ecommerce_Backend/models"

type DatabaseRepo interface {

	//Categories
	AllCategories() ([]models.Categories, error)
	CategoryById(id int) (models.Categories, error)

	//Users
	NewUserDB(user models.Users) (string, error)
	UpdateUserDB(user models.Users) (string, error)
	// GetPassword(id int) (models.Users, error)
	AllUsers() ([]models.DisplayUser, error)
	UserById(id int) (models.DisplayUser, error)
	UserByEmailDB(email string) (models.DisplayUser, error)
	DeleteUserDB(id int) (string, error)

	//Google User
	NewGoogleUserDB(googleUser models.GoogleObject) (string, error)
	GoogleAuthenticate(email string) (models.GoogleObject, error)

	//Authentication
	Authenticate(email string) (models.Users, error)

	//Hats
	AllHats() ([]models.Hats, error)
	NewHatsDB(hats models.Hats) (string, error)
}
