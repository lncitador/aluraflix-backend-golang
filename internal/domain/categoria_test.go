package domain

import (
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCategoria(t *testing.T) {
	t.Run("should create a new categoria", func(t *testing.T) {
		id, _ := vo.NewUniqueEntityID(nil)
		name := "my category"
		color := "#000000"
		idStr := id.ToString()

		input := CategoriaInput{
			Name:      &name,
			Color:     &color,
			UsuarioID: &idStr,
		}

		categoria, err := NewCategoria(input)

		assert.Nil(t, err)
		assert.NotNil(t, categoria)
	})

	t.Run("should return an error when name is empty", func(t *testing.T) {
		name := ""
		color := "#000000"

		input := CategoriaInput{
			Name:  &name,
			Color: &color,
		}

		categoria, err := NewCategoria(input)

		assert.NotNil(t, err)
		assert.Nil(t, categoria)
	})

	t.Run("should return an error when color is empty", func(t *testing.T) {
		name := "my category"
		color := ""

		input := CategoriaInput{
			Name:  &name,
			Color: &color,
		}

		categoria, err := NewCategoria(input)

		assert.NotNil(t, err)
		assert.Nil(t, categoria)
	})

	t.Run("should return an error when color is invalid", func(t *testing.T) {
		name := "my category"
		color := "invalid"

		input := CategoriaInput{
			Name:  &name,
			Color: &color,
		}

		categoria, err := NewCategoria(input)

		assert.NotNil(t, err)
		assert.Nil(t, categoria)
	})
}
