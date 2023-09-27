package inmemory

import (
	"errors"
	e "github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

type VideoRepository struct {
	db       []domain.Video
	relation *CategoriaRepository
}

func NewVideoRepository(relation *CategoriaRepository) *VideoRepository {
	var db []domain.Video
	return &VideoRepository{db, relation}
}

func (r *VideoRepository) FindAll(query *domain.VideoQuery) ([]domain.Video, error) {
	var videos []domain.Video
	if search := query.Search(); search != nil {
		for _, video := range r.db {
			if video.Title == *search || video.Description == *search {
				videos = append(videos, video)
			}
		}
	} else {
		videos = r.db
	}

	if page := query.Page(); page != nil {
		var videosPage []domain.Video
		limit := *query.Limit()
		total := len(videos)
		query.SetTotal(int64(total))

		for i := (*page - 1) * limit; i < *page*limit; i++ {
			if i < total {
				videosPage = append(videosPage, videos[i])

				if len(videosPage) >= len(videos) {
					break
				}
			}
		}

		return videosPage, nil
	}

	return r.db, nil
}

func (r *VideoRepository) FindById(id *v.UniqueEntityID) (*domain.Video, error) {
	if id == nil {
		return nil, errors.New(e.ErrVideoIdIsNull)
	}

	for _, video := range r.db {
		if video.ID.Equals(id) {
			return &video, nil
		}
	}

	return nil, errors.New(e.ErrFindByIdVideo)
}

func (r *VideoRepository) Create(data domain.Video) error {
	if _, err := r.FindById(data.ID); err == nil {
		return errors.New(e.ErrVideoAlreadyExists)
	}

	if _, err := r.relation.FindById(data.CategoryID); err != nil {
		return errors.New(e.ErrCategoriaNotFound)
	}

	r.db = append(r.db, data)

	return nil
}

func (r *VideoRepository) Update(data domain.Video) error {
	for i, video := range r.db {
		if video.ID.Equals(data.ID) {
			r.db[i] = data

			return nil
		}
	}

	return errors.New(e.ErrUpdateVideo)
}

func (r *VideoRepository) Delete(id *v.UniqueEntityID) error {
	for i, video := range r.db {
		if video.ID.Equals(id) {
			r.db = append(r.db[:i], r.db[i+1:]...)

			return nil
		}
	}

	return errors.New(e.ErrDeleteVideo)
}
