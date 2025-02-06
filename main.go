package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"demo/src/core"
	"demo/src/users/infraestructure" // ✅ Agregar esta línea para importar usersInfra correctamente
)

func main() {
	// Inicializar la base de datos
	core.InitDB()

	// Crear router de Gin
	r := gin.Default()

	// Configurar rutas de usuarios
	usersRouter := infrastructure.NewRouter(r) // ✅ Cambiar usersInfra por infrastructure
	usersRouter.Run()

	// Iniciar el servidor en el puerto 8000
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
