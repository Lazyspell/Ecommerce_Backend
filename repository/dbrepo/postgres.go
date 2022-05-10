package dbrepo

import (
	"context"
	"time"

	"github.com/lazyspell/Ecommerce_Backend/models"
)

func (m *postgresDBRepo) AllCategories() ([]models.Categories, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var categories []models.Categories

	query := `select id, title, image_url from categories order by id`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var cat models.Categories
		err := rows.Scan(
			&cat.Id,
			&cat.Title,
			&cat.ImageUrl,
		)
		if err != nil {
			return categories, nil
		}
		categories = append(categories, cat)
	}
	if err = rows.Err(); err != nil {
		return categories, err
	}

	return categories, nil
}

func (m *postgresDBRepo) CategoryById(id int) (models.Categories, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var category models.Categories

	query := `select id, title, image_url from categories where id = $1`

	item := m.DB.QueryRowContext(ctx, query, id)

	err := item.Scan(
		&category.Id,
		&category.ImageUrl,
		&category.Title,
	)
	if err != nil {
		return category, nil
	}

	return category, nil

}
