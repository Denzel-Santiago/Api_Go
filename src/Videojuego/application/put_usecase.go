// application/updateVideojuego_usecase.go
package application

import (
	"demo/src/videojuego/domain/entities"
	"demo/src/videojuego/infraestructure/repositories"
	"fmt"
)

type UpdateVideojuegoUsecase struct {
	Repo repositories.VideojuegoRepository
}

func NewUpdateVideojuegoUsecase(repo repositories.VideojuegoRepository) *UpdateVideojuegoUsecase {
	return &UpdateVideojuegoUsecase{Repo: repo}
}

func (uc *UpdateVideojuegoUsecase) UpdateVideojuego(videojuego *entities.Videojuego) error {
	existingVideojuego, err := uc.Repo.GetByID(videojuego.ID)
	if err != nil {
		return err
	}
	if existingVideojuego == nil {
		return fmt.Errorf("videojuego con ID %d no encontrado", videojuego.ID)
	}
	return uc.Repo.Update(videojuego)
}
