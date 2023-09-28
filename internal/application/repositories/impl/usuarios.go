package impl

import (
	"errors"
	e "github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	. "github.com/lncitador/alura-flix-backend/pkg/errors"
	"gorm.io/gorm"
	"net/http"
)

type UsuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{
		db: db,
	}
}

func (u UsuarioRepository) FindAll(query *struct{}) ([]domain.Usuario, *Error) {
	var usuarios []domain.Usuario

	if err := u.db.Find(&usuarios).Error; err != nil {
		return nil, NewErrorByBadRequest(err)
	}

	return usuarios, nil
}

func (u UsuarioRepository) FindById(id *vo.UniqueEntityID) (*domain.Usuario, *Error) {
	if id == nil {
		return nil, NewError(http.StatusBadRequest, e.ErrUsuarioIdIsNull, "")
	}

	var usuario domain.Usuario
	if err := u.db.First(&usuario, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NewErrorByNotFound(err)
		}

		return nil, NewErrorByBadRequest(err)
	}

	return &usuario, nil
}

func (u UsuarioRepository) Create(data domain.Usuario) *Error {
	if err := u.db.Create(&data).Error; err != nil {
		return NewErrorByBadRequest(err)
	}

	return nil
}

func (u UsuarioRepository) Update(data domain.Usuario) *Error {
	if err := u.db.Save(&data).Error; err != nil {
		return NewErrorByBadRequest(err)
	}

	return nil
}

func (u UsuarioRepository) Delete(id *vo.UniqueEntityID) *Error {
	if id == nil {
		return NewError(http.StatusBadRequest, e.ErrUsuarioIdIsNull, "")
	}

	var usuario domain.Usuario
	if err := u.db.Delete(&usuario, id).Error; err != nil {
		return NewErrorByBadRequest(err)
	}

	return nil
}

func (u UsuarioRepository) FindByEmail(email *string) (*domain.Usuario, *Error) {
	if email == nil {
		return nil, NewError(http.StatusBadRequest, e.ErrUsuarioEmailIsNull, "")
	}

	var usuario domain.Usuario
	if err := u.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NewError(http.StatusNotFound, e.ErrUsuarioNotFound, "")
		}

		return nil, NewErrorByBadRequest(err)
	}

	return &usuario, nil
}
