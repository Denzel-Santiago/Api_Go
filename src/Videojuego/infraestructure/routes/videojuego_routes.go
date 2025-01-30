package routes

import (
	"demo/src/Videojuego/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, controller *controllers.VideojuegoController) {
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("/", controller.CreateVideojuego)
		productRoutes.GET("/:id", controller.GetVideojuego)
		productRoutes.PUT("/:id", controller.UpdateVideojuego)
		productRoutes.DELETE("/:id", controller.DeleteVideojuego)
	}
}
