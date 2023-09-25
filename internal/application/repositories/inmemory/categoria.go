package inmemory

import (
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

type CategoriaRepository struct {
	db []domain.Categoria
}

func NewCategoriaRepository() *CategoriaRepository {
	return &CategoriaRepository{}
}

func (r *CategoriaRepository) FindAll() ([]domain.Categoria, error) {
	return r.db, nil
}

func (r *CategoriaRepository) FindById(id *vo.UniqueEntityID) (*domain.Categoria, error) {
	for _, categoria := range r.db {
		if categoria.ID.Equals(id) {
			return &categoria, nil
		}
	}

	return nil, nil
}

func (r *CategoriaRepository) Create(data domain.Categoria) error {
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

	return nil
}

func (r *CategoriaRepository) Delete(id *vo.UniqueEntityID) error {
	for i, categoria := range r.db {
		if categoria.ID.Equals(id) {
			r.db = append(r.db[:i], r.db[i+1:]...)

			return nil
		}
	}

	return nil
}
