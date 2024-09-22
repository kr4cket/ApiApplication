package repository

import (
	"ApiApplication/pkg/models"

	"github.com/jmoiron/sqlx"
)

type User interface {
	AddUser(user models.User) (bool, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
