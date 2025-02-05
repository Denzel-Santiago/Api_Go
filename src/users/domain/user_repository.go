package domain

import (
	"demo/src/users/domain/entities"
)

type IUser interface {
	Save(user entities.User) error
	Update(id int, user entities.User) error
	Delete(id int) error
	GetAll() ([]entities.User, error)
	FindByID(id int) (entities.User, error) // 🔹 Agregando la función faltante
}
