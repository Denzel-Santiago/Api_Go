package application

import (
	"demo/src/users/domain"
	"demo/src/users/domain/entities"
)

type CreateUserUseCase struct {
	db domain.IUser
}

func NewCreateUserUseCase(db domain.IUser) *CreateUserUseCase {
	return &CreateUserUseCase{
		db: db,
	}
}

func (us *CreateUserUseCase) Run(user *entities.User) (*entities.User, error) {
	err := us.db.Save(*user)
	return user, err
}
