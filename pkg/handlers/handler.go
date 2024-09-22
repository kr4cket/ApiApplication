package handlers

import (
	"ApiApplication/pkg/services"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *services.Service
}

func New(service *services.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/:id", h.getUser)
			users.GET("/", h.getUsers)
			users.POST("/add", h.addUser)
		}

	}

	return router
}
