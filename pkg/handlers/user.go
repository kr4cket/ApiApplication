package handlers

import (
	"ApiApplication/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary AddUser
// @Description Создает пользователя в базе данных
// @ID add-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "Информация о пользователе"
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/users/add [post]
func (h *Handler) addUser(c *gin.Context) {

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Can't parse JSON",
			"errorMessage": err.Error(),
		})
		return
	}

	isAdded, err := h.services.AddUser(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"userAdded":    isAdded,
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"userAdded": isAdded,
	})
}

// @Summary Get User
// @Description Получает пользователя из базы данных
// @ID add-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "Информация о пользователе"
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/users/add [post]
func (h *Handler) getUser(c *gin.Context) {

}

// @Summary Get Users
// @Description Получает всех пользователей
// @ID add-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "Информация о пользователе"
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/users/add [post]
func (h *Handler) getUsers(c *gin.Context) {

}
