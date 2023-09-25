package domain

import (
	"github.com/go-playground/validator/v10"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"time"
)

type Categoria struct {
	Base
	Name  string `gorm:"type:varchar(255);not null"`
	Color string `gorm:"type:varchar(6);not null"`
}

type CategoriaDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CategoriaInput struct {
	Name  *string `json:"nome" validate:"required,min=8,max=255"`
	Color *string `json:"cor" validate:"required,min=6,max=6,hexcolor"`
}

func (i CategoriaInput) validate() error {
	if err := validate.Struct(i); err != nil {
		return err
	}

	return nil
}

// NewCategoria creates a new Categoria instance
func NewCategoria(input CategoriaInput) (*Categoria, error) {
	categoria := Categoria{}
	categoria.prepare()

	if err := input.validate(); err != nil {
		return nil, err
	}

	categoria.Name = *input.Name
	categoria.Color = *input.Color

	return &categoria, nil
}

// MapTo maps Categoria to CategoriaDto struct
func (c *Categoria) MapTo() *CategoriaDto {
	return &CategoriaDto{
		ID:        c.ID.ToString(),
		Name:      c.Name,
		Color:     c.Color,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

// Fill updates the Categoria instance
func (c *Categoria) Fill(input CategoriaInput) error {
	err := validate.Struct(input)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				continue
			}
			return err
		}
	}

	if input.Name != nil {
		c.Name = *input.Name
	}

	if input.Color != nil {
		c.Color = *input.Color
	}

	return nil
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

	return &Categoria{base, d.Name, d.Color}, nil
}
