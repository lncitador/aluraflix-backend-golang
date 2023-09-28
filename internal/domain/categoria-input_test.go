package domain

import (
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoriaInput_prepare(t *testing.T) {
	t.Run("should prepare categoria input", func(t *testing.T) {
		name := "My Category"
		color := "#000000"

		input := CategoriaInput{
			Name:  &name,
			Color: &color,
		}

		input.prepare()

		assert.Equal(t, "my category", *input.Name)
	})

	t.Run("should prepare categoria with non normalized name", func(t *testing.T) {
		name := " My Category   "
		color := "#000000"

		input := CategoriaInput{
			Name:  &name,
			Color: &color,
		}

		input.prepare()

		assert.Equal(t, "my category", *input.Name)
	})
}

func TestCategoriaInput_validate(t *testing.T) {
	t.Run("should validate categoria input", func(t *testing.T) {
		id, _ := vo.NewUniqueEntityID(nil)
		name := "My Category"
		color := "#000000"

		idStr := id.ToString()

		input := CategoriaInput{
			Name:      &name,
			Color:     &color,
			UsuarioID: &idStr,
		}

		err := input.validate()

		assert.Nil(t, err)
	})

	t.Run("should return an error when name is empty", func(t *testing.T) {
		name := ""
		color := "#000000"

		input := CategoriaInput{
			Name:  &name,
			Color: &color,
		}

		err := input.validate()

		assert.NotNil(t, err)
	})

	t.Run("should return an error when name is too short", func(t *testing.T) {
		name := "short"
		color := "#000000"

		input := CategoriaInput{
			Name:  &name,
			Color: &color,
		}

		err := input.validate()

		assert.NotNil(t, err)
	})

	t.Run("should return an error when name is too long", func(t *testing.T) {
		str256 := [256]byte{}
		nameLong := string(str256[:])
		color := "#000000"

		input := CategoriaInput{
			Name:  &nameLong,
			Color: &color,
		}

		err := input.validate()

		assert.NotNil(t, err)
	})

	t.Run("should return an error when color is empty", func(t *testing.T) {
		name := "My Category"
		color := ""

		input := CategoriaInput{
			Name:  &name,
			Color: &color,
		}

		err := input.validate()

		assert.NotNil(t, err)
	})

	t.Run("should return an error when color is too short", func(t *testing.T) {
		name := "My Category"
		color := "#00000"

		input := CategoriaInput{
			Name:  &name,
			Color: &color,
		}

		err := input.validate()

		assert.NotNil(t, err)
	})

	t.Run("should return an error when color is too long", func(t *testing.T) {
		name := "My Category"
		color := "#0000000"

		input := CategoriaInput{
			Name:  &name,
			Color: &color,
		}

		err := input.validate()

		assert.NotNil(t, err)
	})

	t.Run("should return an error when color is invalid", func(t *testing.T) {
		name := "My Category"
		color := "invalid"

		input := CategoriaInput{
			Name:  &name,
			Color: &color,
		}

		err := input.validate()

		assert.NotNil(t, err)
	})
}
