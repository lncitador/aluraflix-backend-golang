package domain

import (
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	. "github.com/lncitador/alura-flix-backend/pkg/errors"
	"net/http"
)

type CategoriaQuery struct {
	usuarioId *vo.UniqueEntityID
}

func (q *CategoriaQuery) UsuarioID() *vo.UniqueEntityID {
	return q.usuarioId
}

func (q *CategoriaQuery) SetUsuarioID(value string) Error {
	if value != "" {
		id, err := vo.NewUniqueEntityID(&value)
		if err != nil {
			return err
		}

		q.usuarioId = id

		return nil
	}

	return NewError(http.StatusBadRequest, "Usuário não informado", "")
}
