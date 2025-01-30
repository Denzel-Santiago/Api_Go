package routes

import (
	"database/sql"
	"net/http"

	"demo/src/Videojuego/application"
	"demo/src/Videojuego/infraestructure/controllers"
	"demo/src/Videojuego/infraestructure/repositories"
	productRoutes "demo/src/Videojuego/infraestructure/routes"
	userRoutes "demo/src/users/infraestructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) http.Handler {
	// Crear router principal con `mux`
	mainRouter := mux.NewRouter()

	// Registrar rutas de usuarios con `mux`
	userRoutes.RegisterUserRoutes(mainRouter, db)

	// Crear router de `gin` para productos
	ginRouter := gin.Default() // Se usa `Default()` para logs y recovery
	VideojuegoRepo := repositories.NewVideojuegoRepository(db)
	VideojuegoUsecase := application.NewManageVideojuegosUsecase(VideojuegoRepo)
	getVideojuegoUsecase := application.NewGetVideojuegoUsecase(VideojuegoRepo)
	updateVideojuegoUsecase := application.NewUpdateVideojuegoUsecase(VideojuegoRepo)
	deleteVideojuegoUsecase := application.NewDeleteVideojuegoUsecase(VideojuegoRepo)

	productController := controllers.NewVideojuegoController(
		VideojuegoUsecase,
		getVideojuegoUsecase,
		updateVideojuegoUsecase,
		deleteVideojuegoUsecase,
	)

	// Registrar rutas de productos en `gin`
	productRoutes.RegisterProductRoutes(ginRouter, productController)

	// Adaptar `ginRouter` para `mux`
	mainRouter.PathPrefix("/products").Handler(http.StripPrefix("/products", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ginRouter.ServeHTTP(w, r)
	})))

	return mainRouter
}
