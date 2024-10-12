package services

import (
	"ApiApplication/pkg/models"
	"errors"
	"github.com/google/uuid"
	"time"
)

var TaskStatuses = map[int]string{
	1: "Started",
	2: "Completed",
	3: "In Progress",
}

var TaskPriority = map[int]string{
	1: "High",
	2: "Middle",
	3: "Low",
}

var tasks = []models.Task{
	{
		Id:          uuid.MustParse("110e8400-e29b-41d4-a716-446655440000"),
		Name:        "Create DB connection",
		Description: "This task must for create API environment",
		Priority:    1,
		Status:      1,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now().Add(86400000),
	},
	{
		Id:          uuid.MustParse("220e8400-e29b-41d4-a716-446655440000"),
		Name:        "Create API skeleton",
		Description: "This task must for create API environment and infrastructure",
		Priority:    1,
		Status:      1,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now().Add(2 * 86400000),
	},
	models.Task{
		Id:          uuid.MustParse("330e8400-e29b-41d4-a716-446655440000"),
		Name:        "Create API routes",
		Description: "We need to create 14 routes for user service",
		Priority:    1,
		Status:      1,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now().Add(3 * 86400000),
	},
	models.Task{
		Id:          uuid.MustParse("440e8400-e29b-41d4-a716-446655440000"),
		Name:        "Unit testing",
		Description: "Test our API",
		Priority:    1,
		Status:      1,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now().Add(4 * 86400000),
	},
	{
		Id:          uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
		Name:        "Deployment",
		Description: "DevOps should to deploy our user service in four days",
		Priority:    1,
		Status:      1,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now().Add(5 * 86400000),
	},
}

type TaskService struct {
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (t *TaskService) GetTask(id uuid.UUID) (models.Task, error) {
	for _, task := range tasks {
		if task.Id == id {
			return task, nil
		}
	}

	err := errors.New("Task not found")

	return models.Task{}, err
}
func (t *TaskService) GetTasks() ([]models.Task, error) {
	return tasks, nil
}
func (t *TaskService) CreateTask(newTask models.Task) (bool, error) {
	for _, task := range tasks {
		if newTask.Id == task.Id {
			return false, errors.New("Task already exists")
		}
	}

	return true, nil
}
func (t *TaskService) DeleteTask(id uuid.UUID) (bool, error) {
	for index, task := range tasks {
		if task.Id == id {
			tasks = remove(tasks, index)
			return true, nil
		}
	}

	return false, errors.New("Task not found")
}
func (t *TaskService) UpdateTask(editedTask models.Task) (bool, error) {
	for index, task := range tasks {
		if task.Id == editedTask.Id {
			tasks[index] = editedTask
			return true, nil
		}
	}

	return false, errors.New("Task not found")
}

func remove(slice []models.Task, s int) []models.Task {
	return append(slice[:s], slice[s+1:]...)
}
