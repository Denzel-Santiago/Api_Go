// application/manageVideojuegos_usecase.go
package application

import (
	"demo/src/videojuego/domain/entities"
	"demo/src/videojuego/infraestructure/repositories"
)

type ManageVideojuegosUsecase struct {
	Repo repositories.VideojuegoRepository
}

func NewManageVideojuegosUsecase(repo repositories.VideojuegoRepository) *ManageVideojuegosUsecase {
	return &ManageVideojuegosUsecase{Repo: repo}
}

func (uc *ManageVideojuegosUsecase) CreateVideojuego(videojuego *entities.Videojuego) error {
	return uc.Repo.Create(videojuego)
}
