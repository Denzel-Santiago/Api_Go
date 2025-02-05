package infrastructure

import (
	"demo/src/users/application"
)

func InitUserDependencies() (
	*CreateUserController,
	*ViewUserController,
	*UpdateUserController,
	*DeleteUserController,
	*ViewAllUsersController, // Nuevo controlador
) {
	// Inicializar el repositorio
	repo := NewMysqlUserRepository()

	// Crear casos de uso
	createUseCase := application.NewCreateUserUseCase(repo)
	viewUseCase := application.NewViewUser(repo)
	updateUseCase := application.NewUpdateUser(repo)
	deleteUseCase := application.NewDeleteUserUseCase(repo)
	viewAllUseCase := application.NewViewAllUsers(repo) // Nuevo caso de uso

	// Crear controladores
	createController := NewCreateUserController(createUseCase)
	viewController := NewViewUserController(viewUseCase)
	updateController := NewUpdateUserController(updateUseCase)
	deleteController := NewDeleteUserController(deleteUseCase)
	viewAllController := NewViewAllUsersController(viewAllUseCase) // Nuevo controlador

	return createController, viewController, updateController, deleteController, viewAllController
}
