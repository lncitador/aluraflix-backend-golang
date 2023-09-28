package impl

import (
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{
		db: db,
	}
}

func (u UsuarioRepository) FindAll(query *struct{}) ([]domain.Usuario, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsuarioRepository) FindById(id *vo.UniqueEntityID) (*domain.Usuario, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsuarioRepository) Create(data domain.Usuario) error {
	//TODO implement me
	panic("implement me")
}

func (u UsuarioRepository) Update(data domain.Usuario) error {
	//TODO implement me
	panic("implement me")
}

func (u UsuarioRepository) Delete(id *vo.UniqueEntityID) error {
	//TODO implement me
	panic("implement me")
}

func (u UsuarioRepository) FindByEmail(email *string) (*domain.Usuario, error) {
	//TODO implement me
	panic("implement me")
}
