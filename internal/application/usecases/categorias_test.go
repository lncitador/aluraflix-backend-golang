package usecases

import (
	"github.com/lncitador/alura-flix-backend/internal/application/repositories/inmemory"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CategoriaSut struct {
	repo      *inmemory.CategoriaRepository
	useCase   *CategoriasUseCase
	constants domain.CategoriaInput
}

var invalidColor = "#FFFFFFF"

func setupCategoriaSut() *CategoriaSut {
	repo := inmemory.NewCategoriaRepository()
	useCase := NewCategoriasUseCase(repo)

	nome := "Name da categoria"
	hexColor := "#FFFFFF"

	return &CategoriaSut{
		repo:    repo,
		useCase: useCase,
		constants: domain.CategoriaInput{
			Name:  &nome,
			Color: &hexColor,
		},
	}
}

func TestCategoriasUseCase_Create(t *testing.T) {
	sut := setupCategoriaSut()
	data := sut.constants

	t.Run("should create a categoria", func(t *testing.T) {
		categoria, err := sut.useCase.Create(data)
		assert.Nil(t, err)
		assert.NotNil(t, categoria)

		assert.Equal(t, *data.Name, categoria.Name)
		assert.Equal(t, *data.Color, categoria.Color)
	})

	t.Run("should not create a categoria with invalid name", func(t *testing.T) {
		data := domain.CategoriaInput{
			Name:  nil,
			Color: data.Color,
		}

		categoria, err := sut.useCase.Create(data)
		assert.NotNil(t, err)
		assert.Nil(t, categoria)
	})

	t.Run("should not create a categoria with invalid data", func(t *testing.T) {
		data := domain.CategoriaInput{
			Name:  nil,
			Color: nil,
		}

		categoria, err := sut.useCase.Create(data)
		assert.NotNil(t, err)
		assert.Nil(t, categoria)
	})

	t.Run("should not create a categoria with invalid color", func(t *testing.T) {
		data := domain.CategoriaInput{
			Name:  data.Name,
			Color: &invalidColor,
		}

		categoria, err := sut.useCase.Create(data)
		assert.NotNil(t, err)
		assert.Nil(t, categoria)
	})
}

func TestCategoriasUseCase_Delete(t *testing.T) {
	sut := setupCategoriaSut()
	data := sut.constants

	t.Run("should delete a categoria", func(t *testing.T) {
		categoria, err := sut.useCase.Create(data)

		id, err := vo.NewUniqueEntityID(&categoria.ID)
		assert.Nil(t, err)
		assert.NotNil(t, categoria)

		err = sut.useCase.Delete(id)
		assert.Nil(t, err)
	})

	t.Run("should not delete a categoria not found", func(t *testing.T) {
		id, _ := vo.NewUniqueEntityID(nil)

		err := sut.useCase.Delete(id)
		assert.NotNil(t, err)
	})
}

func TestCategoriasUseCase_FindAll(t *testing.T) {
	sut := setupCategoriaSut()

	t.Run("should find all categorias", func(t *testing.T) {
		categorias, err := sut.useCase.FindAll()
		assert.Nil(t, err)
		assert.NotNil(t, categorias)
	})
}

func TestCategoriasUseCase_FindById(t *testing.T) {
	sut := setupCategoriaSut()
	data := sut.constants

	t.Run("should find a categoria by id", func(t *testing.T) {
		categoria, err := sut.useCase.Create(data)

		id, err := vo.NewUniqueEntityID(&categoria.ID)
		assert.Nil(t, err)
		assert.NotNil(t, categoria)

		categoria, err = sut.useCase.FindById(id)
		assert.Nil(t, err)
		assert.NotNil(t, categoria)
	})

	t.Run("should not find a categoria with invalid id", func(t *testing.T) {
		categoria, err := sut.useCase.FindById(nil)
		assert.NotNil(t, err)
		assert.Nil(t, categoria)
	})
}

func TestCategoriasUseCase_Update(t *testing.T) {
	sut := setupCategoriaSut()
	data := sut.constants

	t.Run("should update a categoria", func(t *testing.T) {
		categoria, err := sut.useCase.Create(data)

		id, err := vo.NewUniqueEntityID(&categoria.ID)
		assert.Nil(t, err)
		assert.NotNil(t, categoria)

		name := "Name da categoria atualizado"

		data := domain.CategoriaInput{
			Name: &name,
		}

		categoria, err = sut.useCase.Update(id, data)
		assert.Nil(t, err)
		assert.NotNil(t, categoria)
	})

	t.Run("should not update a categoria with invalid id", func(t *testing.T) {
		categoria, err := sut.useCase.Update(nil, data)
		assert.NotNil(t, err)
		assert.Nil(t, categoria)
	})

	t.Run("should not update a categoria with invalid data", func(t *testing.T) {
		categoria, err := sut.useCase.Update(nil, domain.CategoriaInput{})
		assert.NotNil(t, err)
		assert.Nil(t, categoria)
	})

	t.Run("should not update a categoria with invalid color", func(t *testing.T) {
		categoria, err := sut.useCase.Update(nil, domain.CategoriaInput{
			Color: &invalidColor,
		})
		assert.NotNil(t, err)
		assert.Nil(t, categoria)
	})
}
