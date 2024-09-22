package repository

import (
	"ApiApplication/pkg/models"
	"crypto/sha1"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash)
}

func (r *UserPostgres) AddUser(user models.User) (bool, error) {

	pass := generatePasswordHash(user.Password)

	_, err := r.db.Exec("INSERT INTO users (login, password) VALUES ($1, $2)", user.Login, pass)

	if err != nil {
		return false, err
	}

	return true, nil
}
