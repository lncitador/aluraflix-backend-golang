package repositories

import "github.com/lncitador/alura-flix-backend/internal/domain"

type UsuarioRepositoryContract interface {
	RepositoryContract[domain.Usuario, *struct{}]
	FindByEmail(email *string) (*domain.Usuario, error)
}
