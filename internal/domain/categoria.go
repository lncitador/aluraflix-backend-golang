package domain

import (
	"github.com/go-playground/validator/v10"
)

type Categoria struct {
	Base
	Name  string `gorm:"type:varchar(255);not null"`
	Color string `gorm:"type:varchar(7);not null"`
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
