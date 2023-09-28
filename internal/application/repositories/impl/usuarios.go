package impl

import (
	"errors"
	e "github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
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
	var usuarios []domain.Usuario

	if err := u.db.Find(&usuarios).Error; err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (u UsuarioRepository) FindById(id *vo.UniqueEntityID) (*domain.Usuario, error) {
	if id == nil {
		return nil, errors.New(e.ErrUsuarioIdIsNull)
	}

	var usuario domain.Usuario
	if err := u.db.First(&usuario, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(e.ErrFindByIdUsuario)
		}

		return nil, err
	}

	return &usuario, nil
}

func (u UsuarioRepository) Create(data domain.Usuario) error {
	if err := u.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (u UsuarioRepository) Update(data domain.Usuario) error {
	if err := u.db.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (u UsuarioRepository) Delete(id *vo.UniqueEntityID) error {
	if id == nil {
		return errors.New(e.ErrUsuarioIdIsNull)
	}

	var usuario domain.Usuario
	if err := u.db.Delete(&usuario, id).Error; err != nil {
		return err
	}

	return nil
}

func (u UsuarioRepository) FindByEmail(email *string) (*domain.Usuario, error) {
	if email == nil {
		return nil, errors.New(e.ErrUsuarioEmailIsNull)
	}

	var usuario domain.Usuario
	if err := u.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(e.ErrFindByIdUsuario)
		}

		return nil, err
	}

	return &usuario, nil
}
