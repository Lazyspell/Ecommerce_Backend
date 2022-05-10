package dbrepo

import (
	"context"
	"time"

	"github.com/lazyspell/Ecommerce_Backend/models"
)

func (m *postgresDBRepo) Authenticate(email string) (models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.Users

	query := `select id, first_name, last_name, email, password from users where email = $1`

	person := m.DB.QueryRowContext(ctx, query, email)

	err := person.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
