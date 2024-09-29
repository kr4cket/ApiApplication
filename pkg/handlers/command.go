package handlers

import (
	"github.com/gin-gonic/gin"
)

// @Summary CreateModel
// @Description Создает модель на отправленных пользователем данных
// @ID add-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "Информация о пользователе"
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/commands [post]
func (h *Handler) sissoppModel(c *gin.Context) {
	println("Hello World")
}
