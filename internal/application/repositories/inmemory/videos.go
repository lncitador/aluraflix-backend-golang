package inmemory

import (
	"errors"
	e "github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

type VideoRepository struct {
	db []domain.Video
}

func NewVideoRepository() *VideoRepository {
	return &VideoRepository{}
}

func (r *VideoRepository) FindAll() ([]domain.Video, error) {
	return r.db, nil
}

func (r *VideoRepository) FindById(id *v.UniqueEntityID) (*domain.Video, error) {
	for _, video := range r.db {
		if video.ID.Equals(id) {
			return &video, nil
		}
	}

	return nil, errors.New(e.ErrFindByIdVideo)
}

func (r *VideoRepository) Create(data domain.Video) error {
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
