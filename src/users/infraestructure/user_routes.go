package infrastructure

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
	CreateUserController, UpdateUserController, DeleteUserController, ViewAllUsersController := InitUserDependencies()

	userGroup := router.engine.Group("/user")
	{
		userGroup.POST("/", CreateUserController.Run)
		userGroup.PUT("/:id", UpdateUserController.Execute)
		userGroup.DELETE("/:id", DeleteUserController.Run)
		userGroup.GET("/", ViewAllUsersController.Execute)
	}
}
