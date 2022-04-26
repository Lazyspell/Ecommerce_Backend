package repository

import "github.com/lazyspell/Ecommerce_Backend/models"

type DatabaseRepo interface {
	AllCategories() ([]models.Categories, error)
}
