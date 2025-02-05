package infrastructure

import (
	"demo/src/users/application"
)

func InitUserDependencies() (
	*CreateUserController,
	*UpdateUserController,
	*DeleteUserController,
	*ViewAllUsersController,
) {
	repo := NewMysqlUsuarioRepository() // ⚠️ Debe implementar completamente IUser

	createUseCase := application.NewCreateUserUseCase(repo) // ✅ Llamar a NewCreateUserUseCase correctamente
	viewallUseCase := application.NewViewUser(repo)
	deleteUseCase := application.NewDeleteUserUseCase(repo)
	updateUseCase := application.NewUpdateUser(repo)

	CreateController := NewCreateUserController(createUseCase)
	ViewAllController := NewViewAllUsersController(viewallUseCase)
	DeleteController := NewDeleteUserController(deleteUseCase)
	UpdateController := NewUpdateUserController(updateUseCase)

	return CreateController, UpdateController, DeleteController, ViewAllController
}
