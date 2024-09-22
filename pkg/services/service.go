package services

import (
	"ApiApplication/pkg/models"
	"ApiApplication/pkg/repository"
)

type User interface {
	AddUser(user models.User) (bool, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
