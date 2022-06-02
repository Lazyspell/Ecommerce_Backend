package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/lazyspell/Ecommerce_Backend/models"
)

func (m *postgresDBRepo) AllProducts() ([]models.Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var products []models.Products

	query := `select id, product_name, image_url, price, product, created_at, updated_at from products`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Products
		err := rows.Scan(
			&product.Id,
			&product.ProductName,
			&product.ImageUrl,
			&product.Price,
			&product.Product,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return products, nil
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {

		return products, err
	}
	return products, nil

}

func (m *postgresDBRepo) NewProductDB(product models.Products) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into products (product_name, image_url, price, product, created_at, updated_at) values ($1, $2, $3, $4, $5, $6)`

	_, err := m.DB.ExecContext(ctx, query, product.ProductName, product.ImageUrl, product.Price, product.Product, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		log.Println(err)
		return "failed", err
	}
	return "success", nil
}
