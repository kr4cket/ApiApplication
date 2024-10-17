package services

import (
	"ApiApplication/pkg/models"
	"errors"
	"github.com/google/uuid"
)

var projects = []models.Project{}

type ProjectService struct {
}

func newProjectService() *ProjectService {
	return new(ProjectService)
}

func getProjects() ([]models.Project, error) {
	return projects, nil
}

func getTaskByProject(uuid uuid.UUID) ([]models.Task, error) {
	result := []models.Task{}

	for _, task := range tasks {
		if task.ProjectId == uuid {
			result = append(result, task)
		}
	}

	return result, nil
}

func addTaskToProject(projectId uuid.UUID, taskId uuid.UUID) error {
	var projId uuid.UUID

	for _, project := range projects {
		if project.Id == projectId {
			projId = project.Id
		}
	}

	if projId == uuid.Nil {
		return errors.New("Project does not exist")
	}

	for _, task := range tasks {
		if task.Id == taskId && task.ProjectId != uuid.Nil {
			return errors.New("Task already had project")
		}
		if task.Id == taskId && task.ProjectId == uuid.Nil {
			task.ProjectId = projectId
			return nil
		}
	}

	return errors.New("Task Not Found")
}
