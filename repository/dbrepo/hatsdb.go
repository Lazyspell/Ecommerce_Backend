package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/lazyspell/Ecommerce_Backend/models"
)

func (m *postgresDBRepo) AllHats() ([]models.Hats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var hats []models.Hats

	query := `select id, hat_name, image_url, price from hats`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return hats, err
	}

	defer rows.Close()

	for rows.Next() {
		var hat models.Hats
		err := rows.Scan(
			&hat.Id,
			&hat.HatName,
			&hat.ImageUrl,
			&hat.Price,
		)
		if err != nil {
			return hats, nil
		}
		hats = append(hats, hat)
	}

	if err = rows.Err(); err != nil {
		return hats, err
	}
	return hats, nil

}

func (m *postgresDBRepo) NewHatsDB(hats models.Hats) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into hats (hat_name, image_url, price, created_at, updated_at) values ($1, $2, $3, $4, $5)`

	_, err := m.DB.ExecContext(ctx, query, hats.HatName, hats.ImageUrl, hats.Price, hats.CreatedAt, hats.UpdatedAt)
	if err != nil {
		log.Println(err)
		return "failed", err
	}
	return "success", nil
}
