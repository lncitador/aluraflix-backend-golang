package repositories

import (
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	. "github.com/lncitador/alura-flix-backend/pkg/errors"
)

type RepositoryContract[Model any, Query any] interface {
	FindAll(query Query) ([]Model, Error)
	FindById(id *vo.UniqueEntityID) (*Model, Error)
	Create(data Model) Error
	Update(data Model) Error
	Delete(id *vo.UniqueEntityID) Error
}
