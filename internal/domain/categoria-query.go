package domain

import (
	"fmt"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

type CategoriaQuery struct {
	usuarioId *vo.UniqueEntityID
}

func (q *CategoriaQuery) UsuarioID() *vo.UniqueEntityID {
	return q.usuarioId
}

func (q *CategoriaQuery) SetUsuarioID(value string) error {
	if value != "" {
		id, err := vo.NewUniqueEntityID(&value)
		if err != nil {
			return fmt.Errorf("invalid user id")
		}

		q.usuarioId = id

		return nil
	}

	return fmt.Errorf("user id is required")
}
