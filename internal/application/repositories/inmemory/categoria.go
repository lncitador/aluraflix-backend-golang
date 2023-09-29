package inmemory

import (
	e "github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	. "github.com/lncitador/alura-flix-backend/pkg/errors"
	"net/http"
)

type CategoriaRepository struct {
	db []domain.Categoria
}

func NewCategoriaRepository() *CategoriaRepository {
	return &CategoriaRepository{}
}

func (r *CategoriaRepository) FindAll(query *domain.CategoriaQuery) ([]domain.Categoria, Error) {
	return r.db, nil
}

func (r *CategoriaRepository) FindById(id *vo.UniqueEntityID) (*domain.Categoria, Error) {
	if id == nil {
		return nil, NewError(http.StatusBadRequest, e.ErrCategoriaIdIsNull, "")
	}

	for _, categoria := range r.db {
		if categoria.ID.Equals(id) {
			return &categoria, nil
		}
	}

	return nil, NewError(http.StatusNotFound, e.ErrCategoriaNotFound, "")
}

func (r *CategoriaRepository) Create(data domain.Categoria) Error {
	if _, err := r.FindById(data.ID); err == nil {
		return NewError(http.StatusConflict, e.ErrCategoriaAlreadyExists, "")
	}

	r.db = append(r.db, data)
	return nil
}

func (r *CategoriaRepository) Update(data domain.Categoria) Error {
	for i, categoria := range r.db {
		if categoria.ID.Equals(data.ID) {
			r.db[i] = data

			return nil
		}
	}

	return NewError(http.StatusNotFound, e.ErrCategoriaNotFound, "")
}

func (r *CategoriaRepository) Delete(id *vo.UniqueEntityID) Error {
	for i, categoria := range r.db {
		if categoria.ID.Equals(id) {
			r.db = append(r.db[:i], r.db[i+1:]...)

			return nil
		}
	}

	return NewError(http.StatusNotFound, e.ErrCategoriaNotFound, "")
}
