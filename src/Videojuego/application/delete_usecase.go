// application/deleteVideojuego_usecase.go
package application

import "demo/src/videojuego/infraestructure/repositories"

type DeleteVideojuegoUsecase struct {
	Repo repositories.VideojuegoRepository
}

func NewDeleteVideojuegoUsecase(repo repositories.VideojuegoRepository) *DeleteVideojuegoUsecase {
	return &DeleteVideojuegoUsecase{Repo: repo}
}

func (uc *DeleteVideojuegoUsecase) DeleteVideojuego(id int) error {
	return uc.Repo.Delete(id)
}
