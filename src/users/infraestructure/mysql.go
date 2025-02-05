package infrastructure

import (
	"database/sql"
	"demo/src/core"
	"demo/src/users/domain"
	"demo/src/users/domain/entities"
	"fmt"
	"log"
)

type MysqlUsuario struct {
	conn *sql.DB
}

func NewMysqlUsuarioRepository() domain.IUser {
	conn := core.GetDB()
	return &MysqlUsuario{conn: conn}
}

func (mysql *MysqlUsuario) Save(usuario entities.User) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO Usuarios (Nombre, Correo, Contraseña) VALUES (?, ?, ?)",
		usuario.Name,
		usuario.Email,
		usuario.Password,
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

	usuario.SetID(int(idInserted))
	return nil
}

func (mysql *MysqlUsuario) Update(id int, usuario entities.User) error {
	result, err := mysql.conn.Exec(
		"UPDATE Usuarios SET Nombre = ?, Correo = ?, Contraseña = ? WHERE ID = ?",
		usuario.Name,
		usuario.Email,
		usuario.Password,
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

func (mysql *MysqlUsuario) Delete(id int) error {
	result, err := mysql.conn.Exec("DELETE FROM Usuarios WHERE ID = ?", id)
	if err != nil {
		log.Println("Error al eliminar el usuario:", err)
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

func (mysql *MysqlUsuario) FindByID(id int) (entities.User, error) {
	var usuario entities.User
	row := mysql.conn.QueryRow("SELECT ID, Nombre, Correo, Contraseña FROM Usuarios WHERE ID = ?", id)

	err := row.Scan(
		&usuario.ID,
		&usuario.Name,
		&usuario.Email,
		&usuario.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Usuario no encontrado:", err)
			return entities.User{}, fmt.Errorf("usuario con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el usuario por ID:", err)
		return entities.User{}, err
	}

	return usuario, nil
}

func (mysql *MysqlUsuario) GetAll() ([]entities.User, error) {
	var usuarios []entities.User

	rows, err := mysql.conn.Query("SELECT ID, Nombre, Correo, Contraseña FROM Usuarios")
	if err != nil {
		log.Println("Error al obtener todos los usuarios:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var usuario entities.User
		err := rows.Scan(
			&usuario.ID,
			&usuario.Name,
			&usuario.Email,
			&usuario.Password,
		)
		if err != nil {
			log.Println("Error al escanear usuario:", err)
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return usuarios, nil
}
