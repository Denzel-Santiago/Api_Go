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
	repo := NewMysqlUsuarioRepository()

	createUseCase := application.CreateUserUseCase(repo)
	viewallUseCase := application.ViewUser(repo)
	deleteUseCase := application.DeleteUserUseCase(repo)
	updateUseCase := application.UpdateUser(repo)

	CreateController := NewCreateUserController(createUseCase)
	ViewAllController := NewViewAllUsersController(viewallUseCase)
	DeleteController := NewDeleteUserController(deleteUseCase)
	UpdateController := NewUpdateUserController(updateUseCase)

	return CreateController, UpdateController, DeleteController, ViewAllController
}
