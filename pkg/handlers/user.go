package handlers

import (
	"ApiApplication/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

var users = []models.User{
	{
		1,
		"admin",
		"admin",
		0,
	},
	{
		2,
		"user",
		"user",
		1,
	},
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
}

// @Summary Create User
// @Description Создает пользователя системы
// @ID add-user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param input body models.User true "Информация о пользователе"
// @Success 201 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/users/ [post]
func (h *Handler) createUser(c *gin.Context) {

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Can't parse JSON",
			"errorMessage": err.Error(),
		})
		return
	}

	users = append(users, input)

	//isAdded, err := h.services.AddUser(input)

	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, map[string]interface{}{
	//		"userAdded":    isAdded,
	//		"errorMessage": err.Error(),
	//	})
	//	return
	//}

	c.JSON(http.StatusOK, map[string]interface{}{
		"userAdded": true,
	})
}

// @Summary Get User
// @Description Получает пользователя из базы данных
// @ID get-user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path int true "Идентификатор пользователя"
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/users/{id} [get]
func (h *Handler) getUser(c *gin.Context) {

	user := users[0]

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Id,
	})

	strToken, _ := token.SignedString([]byte("4238j0802h0-0_)(*dsjh"))

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": strToken,
	})
}

// @Summary Get Users
// @Description Получает всех пользователей
// @ID get-users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/users/ [get]
func (h *Handler) getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}
