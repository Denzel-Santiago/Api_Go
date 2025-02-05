package main

import (
	"fmt"

	"demo/src/core"
	"demo/src/users/application"
	"demo/src/users/domain/entities"
	usersInfra "demo/src/users/infrastructure"
)

func main() {

	core.InitDB()

	r := gin.Default()

	usersRouter := usersInfra.NewRouter(r)
	usersRouter.Run()

	equiposRouter := equiposInfra.NewRouter(r)
	equiposRouter.Run()

	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
