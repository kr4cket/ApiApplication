package handlers

import (
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("/api")
	{
		users := api.Group("/commands")
		{
			users.POST("/model", h.createModel)
		}
	}

	return router
}
