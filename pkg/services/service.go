package services

import (
	"ApiApplication/pkg/models"
	"ApiApplication/pkg/repository"
	"github.com/google/uuid"
)

type User interface {
	AddUser(user models.User) (bool, error)
}

type Task interface {
	GetTask(id uuid.UUID) (models.Task, error)
	GetTasks() ([]models.Task, error)
	CreateTask(models.Task) (bool, error)
	DeleteTask(id uuid.UUID) (bool, error)
	UpdateTask(task models.Task) (bool, error)
}

type Service struct {
	User
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		Task: NewTaskService(),
	}
}
