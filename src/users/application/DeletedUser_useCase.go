package application

import (
	"demo/src/users/domain"
	"errors"
)

// DeleteUserUseCase maneja la eliminación de usuarios
type DeleteUserUseCase struct {
	db domain.IUser
}

// NewDeleteUserUseCase inicializa el caso de uso de eliminación de usuario
func NewDeleteUserUseCase(db domain.IUser) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		db: db,
	}
}

// Run ejecuta la eliminación de un usuario por ID
func (us *DeleteUserUseCase) Run(id int) error {
	if id <= 0 {
		return errors.New("ID inválido para eliminación")
	}

	err := us.db.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
