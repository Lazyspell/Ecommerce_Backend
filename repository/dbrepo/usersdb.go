package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/lazyspell/Ecommerce_Backend/models"
)

func (m *postgresDBRepo) NewUserDB(user models.Users) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into users (first_name, last_name, email, password) values ($1, $2, $3, $4)`

	_, err := m.DB.ExecContext(ctx, query, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		log.Println(err)
		return "failed", err
	}
	return "success", err

}

func (m *postgresDBRepo) GetPassword(id int) (models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var hashedpass models.Users

	query := `select password from users where id = $1`

	data := m.DB.QueryRowContext(ctx, query, id)
	err := data.Scan(
		&hashedpass.Password,
	)
	if err != nil {
		log.Println(err)
	}

	return hashedpass, err
}
