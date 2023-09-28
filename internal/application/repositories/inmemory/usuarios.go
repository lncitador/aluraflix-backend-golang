package inmemory

import (
	"errors"
	e "github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

type UsuariosRepository struct {
	db []domain.Usuario
}

func NewUsuariosRepository() *UsuariosRepository {
	return &UsuariosRepository{
		db: []domain.Usuario{},
	}
}

func (r *UsuariosRepository) FindAll(query *struct{}) ([]domain.Usuario, error) {
	return r.db, nil
}

func (r *UsuariosRepository) FindById(id *vo.UniqueEntityID) (*domain.Usuario, error) {
	for _, usuario := range r.db {
		if usuario.ID.Equals(id) {
			return &usuario, nil
		}
	}

	return nil, errors.New(e.ErrUsuarioNotFound)
}

func (r *UsuariosRepository) Create(data domain.Usuario) error {
	r.db = append(r.db, data)
	return nil
}

func (r *UsuariosRepository) Update(data domain.Usuario) error {
	for i, usuario := range r.db {
		if usuario.ID.Equals(data.ID) {
			r.db[i] = data
			return nil
		}
	}

	return errors.New(e.ErrUsuarioNotFound)
}

func (r *UsuariosRepository) Delete(id *vo.UniqueEntityID) error {
	for i, usuario := range r.db {
		if usuario.ID.Equals(id) {
			r.db = append(r.db[:i], r.db[i+1:]...)
			return nil
		}
	}

	return errors.New(e.ErrUsuarioNotFound)
}

func (r *UsuariosRepository) FindByEmail(email *string) (*domain.Usuario, error) {
	for _, usuario := range r.db {
		if usuario.Email == *email {
			return &usuario, nil
		}
	}

	return nil, errors.New(e.ErrUsuarioNotFound)
}
