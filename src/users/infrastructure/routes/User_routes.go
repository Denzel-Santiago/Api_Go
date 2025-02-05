package routes

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		engine: engine,
	}
}

func (router *Router) Run() {
	// Inicializar dependencias
	createController, viewController, updateController, deleteController, viewAllController := InitUserDependencies()

	// Grupo de rutas para empleados
	userGroup := router.engine.Group("/user")
	{
		userGroup.POST("/", createController.Run)
		userGroup.GET("/:id", viewController.Execute)
		userGroup.PUT("/:id", updateController.Execute)
		userGroup.DELETE("/:id", deleteController.Run)
		userGroup.GET("/", viewAllController.Execute)
	}
}
