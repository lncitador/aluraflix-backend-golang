package repositories

import (
	"errors"
	e "github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"gorm.io/gorm"
)

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{db}
}

func (r VideoRepository) FindAll() ([]domain.Video, error) {
	var videos []domain.Video
	if err := r.db.Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}

func (r VideoRepository) FindById(id *v.UniqueEntityID) (*domain.Video, error) {
	var video domain.Video
	if err := r.db.First(&video, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(e.ErrFindByIdVideo)
		}

		return nil, err
	}

	return &video, nil
}

func (r VideoRepository) Create(data domain.Video) error {
	if err := r.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r VideoRepository) Update(data domain.Video) error {
	if err := r.db.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r VideoRepository) Delete(id *v.UniqueEntityID) error {
	if err := r.db.Delete(&domain.Video{}, id).Error; err != nil {
		return err
	}

	return nil
}
