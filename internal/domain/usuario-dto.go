package domain

import (
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"time"
)

type UsuarioDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u UsuarioDto) MapFrom() (*Usuario, error) {
	id, err := vo.NewUniqueEntityID(&u.ID)
	if err != nil {
		return nil, err
	}

	base := Base{
		ID:        id,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	return &Usuario{
		Base:  base,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func UsuariosToDto(usuarios []Usuario) *[]UsuarioDto {
	var usuariosDto []UsuarioDto

	for _, usuario := range usuarios {
		usuariosDto = append(usuariosDto, *usuario.MapTo())
	}

	return &usuariosDto
}
