package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"demo/src/core"
	"demo/src/users/application"
	"demo/src/users/domain/entities"
)

type MysqlUser struct {
	conn *sql.DB
}

func NewMysqlUserRepository() domain.IUser {
	conn := core.GetDB()
	return &MysqlUser{conn: conn}
}

func (mysql *MysqlUser) Save(user entities.User) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO users (name, email, password, deleted) VALUES (?, ?, ?, ?)",
		user.Name,
		user.Email,
		user.Password,
		user.Deleted,
	)
	if err != nil {
		log.Println("Error al guardar el usuario:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	user.SetID(int(idInserted))
	return nil
}

func (mysql *MysqlUser) Update(id int, user entities.User) error {
	result, err := mysql.conn.Exec(
		"UPDATE users SET name = ?, email = ?, password = ?, deleted = ? WHERE id = ?",
		user.Name,
		user.Email,
		user.Password,
		user.Deleted,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar el usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró el usuario con ID:", id)
		return fmt.Errorf("usuario con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MysqlUser) Delete(id int) error {
	_, err := mysql.conn.Exec("UPDATE users SET deleted = true WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar (soft delete) el usuario:", err)
		return err
	}
	return nil
}

func (mysql *MysqlUser) FindByID(id int) (entities.User, error) {
	var user entities.User
	row := mysql.conn.QueryRow("SELECT id, name, email, password, deleted FROM users WHERE id = ?", id)

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Deleted,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Usuario no encontrado:", err)
			return entities.User{}, fmt.Errorf("usuario con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el usuario por ID:", err)
		return entities.User{}, err
	}

	return user, nil
}

func (mysql *MysqlUser) GetAll() ([]entities.User, error) {
	var users []entities.User

	rows, err := mysql.conn.Query("SELECT id, name, email, password, deleted FROM users")
	if err != nil {
		log.Println("Error al obtener todos los usuarios:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Deleted,
		)
		if err != nil {
			log.Println("Error al escanear usuario:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return users, nil
}
