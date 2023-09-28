package impl

import (
	"errors"
	e "github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"gorm.io/gorm"
)

type CategoriaRepository struct {
	db *gorm.DB
}

func NewCategoriaRepository(db *gorm.DB) *CategoriaRepository {
	return &CategoriaRepository{db}
}

func (r CategoriaRepository) FindAll(query *struct{}) ([]domain.Categoria, error) {
	var categorias []domain.Categoria
	if err := r.db.Find(&categorias).Error; err != nil {
		return nil, err
	}

	return categorias, nil
}

func (r CategoriaRepository) FindById(id *vo.UniqueEntityID) (*domain.Categoria, error) {
	if id == nil {
		return nil, errors.New(e.ErrCategoriaIdIsNull)
	}
	var categoria domain.Categoria
	if err := r.db.First(&categoria, id).Error; err != nil {
		return nil, err
	}

	return &categoria, nil
}

func (r CategoriaRepository) Create(data domain.Categoria) error {
	if err := r.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r CategoriaRepository) Update(data domain.Categoria) error {
	if err := r.db.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r CategoriaRepository) Delete(id *vo.UniqueEntityID) error {
	if err := r.db.Delete(&domain.Categoria{}, id).Error; err != nil {
		return err
	}

	return nil
}
