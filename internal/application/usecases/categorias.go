package usecases

import (
	"github.com/lncitador/alura-flix-backend/internal/application/repositories"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

type CategoriaRepositoryContract repositories.RepositoryContract[domain.Categoria, domain.CategoriaQuery]

type CategoriasUseCase struct {
	CategoriaRepositoryContract
}

func NewCategoriasUseCase(contract CategoriaRepositoryContract) *CategoriasUseCase {
	return &CategoriasUseCase{contract}
}

func (c *CategoriasUseCase) FindAll() (*[]domain.CategoriaDto, error) {
	categorias, err := c.CategoriaRepositoryContract.FindAll(domain.CategoriaQuery{})
	if err != nil {
		return nil, err
	}

	return domain.CategoriasToDto(categorias), nil
}

func (c *CategoriasUseCase) FindById(id *vo.UniqueEntityID) (*domain.CategoriaDto, error) {
	categoria, err := c.CategoriaRepositoryContract.FindById(id)
	if err != nil {
		return nil, err
	}

	return categoria.MapTo(), nil
}

func (c *CategoriasUseCase) Create(data domain.CategoriaInput) (*domain.CategoriaDto, error) {
	categoria, err := domain.NewCategoria(data)
	if err != nil {
		return nil, err
	}

	if err := c.CategoriaRepositoryContract.Create(*categoria); err != nil {
		return nil, err
	}

	return categoria.MapTo(), nil
}

func (c *CategoriasUseCase) Update(id *vo.UniqueEntityID, data domain.CategoriaInput) (*domain.CategoriaDto, error) {
	categoria, err := c.CategoriaRepositoryContract.FindById(id)
	if err != nil {
		return nil, err
	}

	if err := categoria.Fill(data); err != nil {
		return nil, err
	}

	if err := c.CategoriaRepositoryContract.Update(*categoria); err != nil {
		return nil, err
	}

	return categoria.MapTo(), nil
}

func (c *CategoriasUseCase) Delete(id *vo.UniqueEntityID) error {
	if err := c.CategoriaRepositoryContract.Delete(id); err != nil {
		return err
	}

	return nil
}
