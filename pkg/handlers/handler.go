package handlers

import (
	"ApiApplication/pkg/services"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "ApiApplication/docs"
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
			users.POST("/", h.createUser)
		}
		tasks := api.Group("/tasks")
		{
			tasks.POST("/", h.addTask)
			tasks.GET("/:id", h.getTask)
			tasks.GET("/", h.getTasks)
			tasks.DELETE("/:id", h.deleteTask)
			tasks.PUT("/:id", h.updateTask)
		}
		projects := api.Group("/projects")
		{
			projects.GET("/:id/tasks", h.getProjectTasksById)
			projects.POST("/:id/tasks", h.addTaskToProject)
			projects.GET("/", h.getProjects)
		}

	}

	return router
}
