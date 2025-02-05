package application

import (
	"demo/src/users/domain"
	"demo/src/users/domain/entities"
)

type ViewUser struct {
	db domain.IUser
}

func NewViewUser(db domain.IUser) *ViewUser {
	return &ViewUser{
		db: db,
	}
}

func (us *ViewUser) Execute() ([]entities.User, error) {
	return us.db.GetAll()
}
