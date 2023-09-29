package impl

import (
	e "github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	. "github.com/lncitador/alura-flix-backend/pkg/errors"
	"gorm.io/gorm"
	"net/http"
)

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{db}
}

func (r VideoRepository) FindAll(query *domain.VideoQuery) ([]domain.Video, Error) {
	var videos []domain.Video

	if usuarioId := query.UsuarioID(); usuarioId != nil {
		r.db = r.db.Where("usuario_id = ?", usuarioId.ToString())
	}

	if search := query.Search(); search != nil {
		r.db = r.db.Where("title LIKE ?", search)
		r.db = r.db.Or("description LIKE ?", search)
	}

	if page := query.Page(); page != nil {
		limit := *query.Limit()
		r.db.Model(&domain.Video{}).Count(query.Total())
		r.db = r.db.Limit(limit).Offset((*page - 1) * limit)
	}

	if err := r.db.Find(&videos).Error; err != nil {
		return nil, NewErrorByBadRequest(err)
	}

	return videos, nil
}

func (r VideoRepository) FindById(id *v.UniqueEntityID) (*domain.Video, Error) {
	if id == nil {
		return nil, NewError(http.StatusBadRequest, e.ErrVideoIdIsNull, "")
	}

	var video domain.Video
	if err := r.db.First(&video, id).Error; err != nil {
		return nil, NewErrorByBadRequest(err)
	}

	return &video, nil
}

func (r VideoRepository) Create(data domain.Video) Error {
	if err := r.db.Create(&data).Error; err != nil {
		return NewErrorByBadRequest(err)
	}

	return nil
}

func (r VideoRepository) Update(data domain.Video) Error {
	if err := r.db.Save(&data).Error; err != nil {
		return NewErrorByBadRequest(err)
	}

	return nil
}

func (r VideoRepository) Delete(id *v.UniqueEntityID) Error {
	if id == nil {
		return NewError(http.StatusBadRequest, e.ErrVideoIdIsNull, "")
	}

	if err := r.db.Delete(&domain.Video{}, id).Error; err != nil {
		return NewErrorByBadRequest(err)
	}

	return nil
}
