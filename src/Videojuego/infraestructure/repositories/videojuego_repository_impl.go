package repositories

import (
	"database/sql"
	"demo/src/videojuego/domain/entities"
	"fmt"
)

type VideojuegoRepository interface {
	Create(videojuego *entities.Videojuego) error
	GetByID(id int) (*entities.Videojuego, error)
	Update(videojuego *entities.Videojuego) error
	Delete(id int) error
}

type VideojuegoRepositoryImpl struct {
	DB *sql.DB
}

func NewVideojuegoRepository(db *sql.DB) VideojuegoRepository {
	return &VideojuegoRepositoryImpl{DB: db}
}

func (r *VideojuegoRepositoryImpl) Create(videojuego *entities.Videojuego) error {
	query := "INSERT INTO videojuegos (nombre, descripcion, precio) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, videojuego.Nombre, videojuego.Descripcion, videojuego.Precio)
	return err
}

func (r *VideojuegoRepositoryImpl) GetByID(id int) (*entities.Videojuego, error) {
	query := "SELECT id, nombre, descripcion, precio FROM videojuegos WHERE id = ?"
	row := r.DB.QueryRow(query, id)
	var videojuego entities.Videojuego
	err := row.Scan(&videojuego.ID, &videojuego.Nombre, &videojuego.Descripcion, &videojuego.Precio)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &videojuego, nil
}

func (r *VideojuegoRepositoryImpl) Update(videojuego *entities.Videojuego) error {
	query := "UPDATE videojuegos SET nombre=?, descripcion=?, precio=? WHERE id=?"
	result, err := r.DB.Exec(query, videojuego.Nombre, videojuego.Descripcion, videojuego.Precio, videojuego.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el videojuego con ID %d", videojuego.ID)
	}

	return nil
}

func (r *VideojuegoRepositoryImpl) Delete(id int) error {
	query := "DELETE FROM videojuegos WHERE id=?"
	_, err := r.DB.Exec(query, id)
	return err
}
