package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/lazyspell/Ecommerce_Backend/models"
)

func (m *postgresDBRepo) AllUsers() ([]models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var users []models.Users

	query := `select id, first_name, last_name, email from users`

	rows, err := m.DB.QueryContext(ctx, query)
	log.Println(rows)

	if err != nil {
		return users, err
	}
	log.Println(2)

	defer rows.Close()

	for rows.Next() {
		log.Println(3)

		var people models.Users
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

func (m *postgresDBRepo) UserById(id int) (models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.Users

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

func (m *postgresDBRepo) DeleteUserDB(id int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `delete from users where id=$1`

	m.DB.QueryRowContext(ctx, query, id)

	log.Println()

	return "success", nil

}
