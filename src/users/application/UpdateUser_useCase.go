package application

import (
	"demo/src/users/domain"
	"demo/src/users/domain/entities"
)

type UpdateUser struct {
	db domain.IUser
}

func NewUpdateUser(db domain.IUser) *UpdateUser {
	return &UpdateUser{
		db: db,
	}
}

func (us *UpdateUser) Execute(id int, user entities.User) error {
	return us.db.Update(id, user)
}
