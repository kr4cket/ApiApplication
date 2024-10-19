package handlers

import (
	"github.com/gin-gonic/gin"
)

// @Summary CreateModel
// @Description Создает модель на основе отправленных пользователем данных
// @ID create-sissop-model
// @Accept  json
// @Produce  json
// @Success 200 {boolean} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/commands/sissop [post]
func (h *Handler) sissoppModel(c *gin.Context) {
}
