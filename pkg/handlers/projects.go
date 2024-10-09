package handlers

import (
	"github.com/gin-gonic/gin"
)

// @Summary Get Task by project
// @Description Получает задачи конкретного проекта
// @ID get-task-by-project
// @Tags Projects
// @Accept  json
// @Produce  json
// @Param id path int true "Идентификатор проекта"
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/projects/{id}/tasks [get]
func (h *Handler) getProjectTasksById(c *gin.Context) {

}

// @Summary Get Projects
// @Description Получение всех существующих проектов
// @ID get-projects
// @Tags Projects
// @Accept  json
// @Produce  json
// @Param input body models.Project true "Информация о проекте"
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/projects/ [get]
func (h *Handler) getProjects(c *gin.Context) {

}

// @Summary Add task to project
// @Description Добавляет задачу в существующий проект
// @ID add-project
// @Tags Projects
// @Accept  json
// @Produce  json
// @Param id path int true "Идентификатор проекта"
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/projects/{id}/tasks [post]
func (h *Handler) addTaskToProject(c *gin.Context) {

}
