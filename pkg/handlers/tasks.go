package handlers

import (
	"ApiApplication/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// @Summary Add Task
// @Description Создает задачу в базе данных
// @ID add-task
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param input body models.Task true "Информация о задаче"
// @Success 201 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/tasks/ [post]
func (h *Handler) addTask(c *gin.Context) {

	var input models.Task

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Can't parse JSON",
			"errorMessage": err.Error(),
		})
		return
	}

	status, err := h.services.CreateTask(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Error creating task",
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"taskAdded": status,
	})
}

// @Summary Get Task
// @Description Получает задачу из базы данных
// @ID get-task
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Идентификатор задачи"
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/tasks/{id} [get]
func (h *Handler) getTask(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Incorrect id",
			"errorMessage": err.Error(),
		})
	}

	task, err := h.services.GetTask(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":        "Error getting task",
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"task": task,
	})
}

// @Summary Get Tasks
// @Description Получает все существующие задачи
// @ID get-tasks
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/tasks/ [get]
func (h *Handler) getTasks(c *gin.Context) {

	tasks, err := h.services.GetTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":        "Error getting tasks",
			"errorMessage": err.Error(),
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})
}

// @Summary Delete Task
// @Description Удаляет выбранную задачу
// @ID delete-task
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Идентификатор задачи"
// @Success 204 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/tasks/{id} [delete]
func (h *Handler) deleteTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Incorrect id",
			"errorMessage": err.Error(),
		})
	}

	status, err := h.services.DeleteTask(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":        "Error deleting task",
			"errorMessage": err.Error(),
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"taskDeleted": status,
	})

}

// @Summary Update Task
// @Description Обновляет информацию о задаче
// @ID update-task
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Идентификатор задачи"
// @Success 204 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/tasks/{id} [put]
func (h *Handler) updateTask(c *gin.Context) {
	var input models.Task
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Can't parse JSON",
			"errorMessage": err.Error(),
		})
	}

	status, err := h.services.UpdateTask(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Error updating task",
			"errorMessage": err.Error(),
		})
	}

	c.JSON(http.StatusNoContent, map[string]interface{}{
		"taskUpdated": status,
	})

}
