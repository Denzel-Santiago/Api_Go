package application

import (
	"demo/src/users/domain"
)

type DeleteUserUseCase struct {
	db domain.IUser
}

func NewDeleteUserUseCase(db domain.IUser) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		db: db,
	}
}

func (us *DeleteUserUseCase) Run(id int) error {
	return us.db.Delete(id)
}
