package dbrepo

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (m *postgresDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int

	var hashedPassword string

	query := `select id, password from users where email =$1`

	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&id,
		&hashedPassword,
	)
	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
