package handlers

import (
	"ApiApplication/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var tasks = []models.Task{
	models.Task{
		Id:          1,
		Name:        "First Task",
		Description: "This task for me",
		Status:      "Ok",
		CreatedAt:   "11.01.2022",
		CompletedAt: "11.01.2022",
	},
	models.Task{
		Id:          2,
		Name:        "Second Task",
		Description: "This task for me",
		Status:      "Ok",
		CreatedAt:   "11.01.2022",
		CompletedAt: "11.01.2022",
	},
	models.Task{
		Id:          3,
		Name:        "Third Task",
		Description: "This task for me",
		Status:      "Ok",
		CreatedAt:   "11.01.2022",
		CompletedAt: "11.01.2022",
	},
	models.Task{
		Id:          4,
		Name:        "Fourth Task",
		Description: "This task for me",
		Status:      "Ok",
		CreatedAt:   "11.01.2022",
		CompletedAt: "11.01.2022",
	},
	models.Task{
		Id:          5,
		Name:        "Fifth Task",
		Description: "This task for me",
		Status:      "Ok",
		CreatedAt:   "11.01.2022",
		CompletedAt: "11.01.2022",
	},
}

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

	tasks = append(tasks, input)
	c.JSON(http.StatusOK, map[string]interface{}{
		"taskAdded": true,
	})
}

// @Summary Get Task
// @Description Получает задачу из базы данных
// @ID get-task
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Идентификатор задачи"
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/tasks/{id} [get]
func (h *Handler) getTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":        "Id not found",
			"errorMessage": err.Error(),
		})
		return
	}

	if len(tasks) <= id {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Id doesn't exist",
		})
		return
	}

	task := tasks[id]

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
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":        "Id not found",
			"errorMessage": err.Error(),
		})
		return
	}

	tasks = remove(tasks, id)

	c.JSON(http.StatusOK, map[string]interface{}{
		"taskDeleted": true,
	})

}

func remove(slice []models.Task, s int) []models.Task {
	return append(slice[:s], slice[s+1:]...)
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

}
