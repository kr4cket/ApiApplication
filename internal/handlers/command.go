package handlers

import (
	"github.com/gin-gonic/gin"
)

// @Summary CreateModel
// @Description Описание модели
// @ID create-smth-model
// @Accept  json
// @Produce  json
// @Success 200 {boolean} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/commands/model [post]
func (h *Handler) createModel(c *gin.Context) {
}
