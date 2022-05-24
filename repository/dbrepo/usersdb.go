package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/lazyspell/Ecommerce_Backend/models"
)

func (m *postgresDBRepo) AllUsers() ([]models.DisplayUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var users []models.DisplayUser

	query := `select id, first_name, last_name, email from users`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {

		var people models.DisplayUser
		err := rows.Scan(
			&people.Id,
			&people.FirstName,
			&people.LastName,
			&people.Email,
		)
		if err != nil {
			return users, nil
		}
		users = append(users, people)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (m *postgresDBRepo) UserById(id int) (models.DisplayUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.DisplayUser

	query := `select id, first_name, last_name, email from users where id = $1`

	person := m.DB.QueryRowContext(ctx, query, id)

	err := person.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	)
	if err != nil {
		return user, nil
	}

	return user, nil

}

func (m *postgresDBRepo) NewUserDB(user models.Users) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into users (first_name, last_name, email, password, created_at, updated_at, "authorization") values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, query, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.Authorization)
	if err != nil {
		log.Println(err)
		return "failed", err
	}
	return "success", nil

}

func (m *postgresDBRepo) NewGoogleUserDB(googleUser models.GoogleObject) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into google_user (email, name, given_name) values ($1, $2, $3)`
	_, err := m.DB.ExecContext(ctx, query, googleUser.Email, googleUser.Name, googleUser.GivenName)
	if err != nil {
		return "failed", err
	}
	return "success", nil

}

func (m *postgresDBRepo) DeleteUserDB(id int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `delete from users where id=$1`

	m.DB.QueryRowContext(ctx, query, id)

	return "success", nil

}
