package repositories

import (
	"github.com/lncitador/alura-flix-backend/internal/domain"
	. "github.com/lncitador/alura-flix-backend/pkg/errors"
)

type UsuarioRepositoryContract interface {
	RepositoryContract[domain.Usuario, *struct{}]
	FindByEmail(email *string) (*domain.Usuario, Error)
}
