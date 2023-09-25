package repositories

import (
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

type RepositoryContract[Model any, Query any] interface {
	FindAll(query Query) ([]Model, error)
	FindById(id *vo.UniqueEntityID) (*Model, error)
	Create(data Model) error
	Update(data Model) error
	Delete(id *vo.UniqueEntityID) error
}
