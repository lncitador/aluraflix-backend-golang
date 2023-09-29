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

type CategoriaRepository struct {
	db *gorm.DB
}

func NewCategoriaRepository(db *gorm.DB) *CategoriaRepository {
	return &CategoriaRepository{db}
}

func (r CategoriaRepository) FindAll(query *domain.CategoriaQuery) ([]domain.Categoria, Error) {
	if usuarioID := query.UsuarioID(); usuarioID != nil {
		r.db = r.db.Where("usuario_id = ?", usuarioID)
	}

	var categorias []domain.Categoria
	if err := r.db.Find(&categorias).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NewErrorByNotFound(err)
		}

		return nil, NewErrorByBadRequest(err)
	}

	return categorias, nil
}

func (r CategoriaRepository) FindById(id *vo.UniqueEntityID) (*domain.Categoria, Error) {
	if id == nil {
		return nil, NewError(http.StatusBadRequest, e.ErrCategoriaIdIsNull, "")
	}
	var categoria domain.Categoria
	if err := r.db.First(&categoria, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NewErrorByNotFound(err)
		}

		return nil, NewErrorByBadRequest(err)
	}

	return &categoria, nil
}

func (r CategoriaRepository) Create(data domain.Categoria) Error {
	if err := r.db.Create(&data).Error; err != nil {
		return NewErrorByBadRequest(err)
	}

	return nil
}

func (r CategoriaRepository) Update(data domain.Categoria) Error {
	if err := r.db.Save(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NewErrorByNotFound(err)
		}

		return NewErrorByBadRequest(err)
	}

	return nil
}

func (r CategoriaRepository) Delete(id *vo.UniqueEntityID) Error {
	if err := r.db.Delete(&domain.Categoria{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NewErrorByNotFound(err)
		}

		return NewErrorByBadRequest(err)
	}

	return nil
}
