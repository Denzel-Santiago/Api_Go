package application

import (
	"demo/src/users/domain"
	"demo/src/users/domain/entities"
	"errors"
)

// UpdateUser maneja la actualización de usuarios
type UpdateUser struct {
	db domain.IUser
}

// NewUpdateUser inicializa el caso de uso para actualizar usuarios
func NewUpdateUser(db domain.IUser) *UpdateUser {
	return &UpdateUser{
		db: db,
	}
}

// Execute actualiza un usuario en la base de datos
func (us *UpdateUser) Execute(id int, user entities.User) error {
	if id <= 0 {
		return errors.New("ID inválido para actualización")
	}

	// Validaciones básicas antes de actualizar
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("los campos Nombre, Email y Contraseña no pueden estar vacíos")
	}

	err := us.db.Update(id, user)
	if err != nil {
		return err
	}

	return nil
}
