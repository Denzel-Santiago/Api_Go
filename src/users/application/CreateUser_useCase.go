package application

import (
	"demo/src/users/domain"
	"demo/src/users/domain/entities"
	"errors"
)

// CreateUserUseCase maneja la creación de usuarios
type CreateUserUseCase struct {
	db domain.IUser
}

// NewCreateUserUseCase inicializa el caso de uso de creación de usuario
func NewCreateUserUseCase(db domain.IUser) *CreateUserUseCase {
	return &CreateUserUseCase{
		db: db,
	}
}

// Run ejecuta la creación de usuario
func (us *CreateUserUseCase) Run(user *entities.User) (*entities.User, error) {
	if user == nil {
		return nil, errors.New("el usuario no puede ser nulo")
	}

	err := us.db.Save(*user) // Se pasa por referencia
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindUserByIDUseCase maneja la búsqueda de usuarios por ID
type FindUserByIDUseCase struct {
	db domain.IUser
}

// NewFindUserByIDUseCase inicializa el caso de uso para encontrar un usuario por ID
func NewFindUserByIDUseCase(db domain.IUser) *FindUserByIDUseCase {
	return &FindUserByIDUseCase{
		db: db,
	}
}

// Run ejecuta la búsqueda de usuario por ID
func (us *FindUserByIDUseCase) Run(id int) (*entities.User, error) {
	if id <= 0 {
		return nil, errors.New("ID inválido")
	}

	user, err := us.db.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
