package application

import (
	"demo/src/users/domain"
	"demo/src/users/domain/entities"
	"errors"
	"log"
)

// ViewUser maneja la visualización de todos los usuarios
type ViewUser struct {
	db domain.IUser
}

// NewViewUser inicializa el caso de uso para obtener usuarios
func NewViewUser(db domain.IUser) *ViewUser {
	return &ViewUser{
		db: db,
	}
}

// Execute obtiene todos los usuarios de la base de datos
func (us *ViewUser) Execute() ([]entities.User, error) {
	users, err := us.db.GetAll()
	if err != nil {
		log.Println("Error al obtener usuarios:", err)
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("no hay usuarios registrados")
	}

	return users, nil
}
