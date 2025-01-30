// application/getVideojuego_usecase.go
package application

import (
	"demo/src/videojuego/domain/entities"
	"demo/src/videojuego/infraestructure/repositories"
)

type GetVideojuegoUsecase struct {
	Repo repositories.VideojuegoRepository
}

func NewGetVideojuegoUsecase(repo repositories.VideojuegoRepository) *GetVideojuegoUsecase {
	return &GetVideojuegoUsecase{Repo: repo}
}

func (uc *GetVideojuegoUsecase) GetVideojuegoByID(id int) (*entities.Videojuego, error) {
	return uc.Repo.GetByID(id)
}
