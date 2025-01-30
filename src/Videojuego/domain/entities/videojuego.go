// domain/entities/videojuego.go
package entities

type Videojuego struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Genero      string  `json:"genero"`
	Plataforma  string  `json:"plataforma"`
	Precio      float64 `json:"precio"`
}
