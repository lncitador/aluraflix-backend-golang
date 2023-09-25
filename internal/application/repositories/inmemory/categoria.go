package inmemory

import (
	"errors"
	e "github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

type CategoriaRepository struct {
	db []domain.Categoria
}

func NewCategoriaRepository() *CategoriaRepository {
	return &CategoriaRepository{}
}

func (r *CategoriaRepository) FindAll(query domain.VideoQuery) ([]domain.Categoria, error) {
	return r.db, nil
}

func (r *CategoriaRepository) FindById(id *vo.UniqueEntityID) (*domain.Categoria, error) {
	if id == nil {
		return nil, errors.New(e.ErrCategoriaIdIsNull)
	}

	for _, categoria := range r.db {
		if categoria.ID.Equals(id) {
			return &categoria, nil
		}
	}

	return nil, errors.New(e.ErrFindByIdCategoria)
}

func (r *CategoriaRepository) Create(data domain.Categoria) error {
	if _, err := r.FindById(data.ID); err == nil {
		return errors.New(e.ErrCategoriaAlreadyExists)
	}

	r.db = append(r.db, data)
	return nil
}

func (r *CategoriaRepository) Update(data domain.Categoria) error {
	for i, categoria := range r.db {
		if categoria.ID.Equals(data.ID) {
			r.db[i] = data

			return nil
		}
	}

	return errors.New(e.ErrUpdateCategoria)
}

func (r *CategoriaRepository) Delete(id *vo.UniqueEntityID) error {
	for i, categoria := range r.db {
		if categoria.ID.Equals(id) {
			r.db = append(r.db[:i], r.db[i+1:]...)

			return nil
		}
	}

	return errors.New(e.ErrDeleteCategoria)
}
