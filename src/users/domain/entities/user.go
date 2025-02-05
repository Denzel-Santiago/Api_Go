package entities

import "time"

type User struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	FechaRegistro time.Time `json:"fecha_registro"`
	TipoUsuario   string    `json:"tipo_usuario"`
}

func NewUser(id int, Name, Email, Password string, FechaRegistro time.Time, TipoUsuario string) *User {
	return &User{
		ID:            id,
		Name:          Name,
		Email:         Email,
		Password:      Password,
		FechaRegistro: FechaRegistro,
		TipoUsuario:   TipoUsuario,
	}
}

func (u *User) GetID() int {
	return u.ID
}

func (u *User) SetID(id int) {
	u.ID = id
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetPassword(password string) {
	u.Password = password
}
