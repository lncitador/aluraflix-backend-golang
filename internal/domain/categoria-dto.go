package domain

import (
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"time"
)

type CategoriaDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	UsuarioID string    `json:"usuarioId"`
	Usuario   *Usuario  `json:"usuario,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CategoriasToDto maps a slice of Categoria to a slice of CategoriaDto
func CategoriasToDto(categorias []Categoria) *[]CategoriaDto {
	var categoriasDto []CategoriaDto

	for _, categoria := range categorias {
		categoriasDto = append(categoriasDto, *categoria.MapTo())
	}

	return &categoriasDto
}

// MapFrom maps CategoriaDto to Categoria struct
func (d CategoriaDto) MapFrom() (*Categoria, error) {
	id, err := vo.NewUniqueEntityID(&d.ID)
	if err != nil {
		return nil, err
	}

	base := Base{
		ID:        id,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}

	usuarioId, err := vo.NewUniqueEntityID(&d.UsuarioID)
	if err != nil {
		return nil, err
	}

	return &Categoria{
		Base:      base,
		Name:      d.Name,
		Color:     d.Color,
		UsuarioID: usuarioId,
	}, nil
}
