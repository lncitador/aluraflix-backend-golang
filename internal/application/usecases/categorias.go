package usecases

import (
	"github.com/lncitador/alura-flix-backend/internal/application/repositories"
	e "github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	. "github.com/lncitador/alura-flix-backend/pkg/errors"
	"net/http"
)

type CategoriasUseCase struct {
	repositories.CategoriaRepositoryContract
}

func NewCategoriasUseCase(contract repositories.CategoriaRepositoryContract) *CategoriasUseCase {
	return &CategoriasUseCase{contract}
}

func (c *CategoriasUseCase) FindAll(query *domain.CategoriaQuery) (*[]domain.CategoriaDto, Error) {
	categorias, err := c.CategoriaRepositoryContract.FindAll(query)
	if err != nil {
		return nil, err
	}

	return domain.CategoriasToDto(categorias), nil
}

func (c *CategoriasUseCase) FindById(id *vo.UniqueEntityID) (*domain.CategoriaDto, Error) {
	categoria, err := c.CategoriaRepositoryContract.FindById(id)
	if err != nil {
		return nil, err
	}

	return categoria.MapTo(), nil
}

func (c *CategoriasUseCase) Create(data domain.CategoriaInput) (*domain.CategoriaDto, Error) {
	categoria, err := domain.NewCategoria(data)
	if err != nil {
		return nil, err
	}

	if err := c.CategoriaRepositoryContract.Create(*categoria); err != nil {
		return nil, err
	}

	return categoria.MapTo(), nil
}

func (c *CategoriasUseCase) Update(id *vo.UniqueEntityID, data domain.CategoriaInput) (*domain.CategoriaDto, Error) {
	categoria, err := c.CategoriaRepositoryContract.FindById(id)
	if err != nil {
		return nil, err
	}

	usuarioId, err := vo.NewUniqueEntityID(data.UsuarioID)
	if err != nil {
		return nil, err
	}

	if !categoria.UsuarioID.Equals(usuarioId) {
		return nil, NewError(http.StatusUnauthorized, e.ErrUsuarioUnauthorized, "")
	}

	if err := categoria.Fill(data); err != nil {
		return nil, err
	}

	if err := c.CategoriaRepositoryContract.Update(*categoria); err != nil {
		return nil, err
	}

	return categoria.MapTo(), nil
}

func (c *CategoriasUseCase) Delete(id *vo.UniqueEntityID) Error {
	if err := c.CategoriaRepositoryContract.Delete(id); err != nil {
		return err
	}

	return nil
}
