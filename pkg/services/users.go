package services

import (
	"ApiApplication/pkg/models"
	"ApiApplication/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddUser(user models.User) (bool, error) {
	return s.repo.AddUser(user)
}
