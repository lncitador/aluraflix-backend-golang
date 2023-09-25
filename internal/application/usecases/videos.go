package usecases

import (
	"fmt"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

type VideosUseCase struct {
	videoRepository *repositories.VideoRepository
}

func NewVideosUseCase(videoRepository *repositories.VideoRepository) *VideosUseCase {
	return &VideosUseCase{videoRepository}
}

func (v VideosUseCase) FindAll() (*[]domain.VideoDto, error) {
	videos, err := v.videoRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return domain.VideosToDto(videos), nil
}

func (v VideosUseCase) FindById(id *vo.UniqueEntityID) (*domain.VideoDto, error) {
	video, err := v.videoRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return video.MapTo(), nil
}

func (v VideosUseCase) Create(data domain.VideoInput) (*domain.VideoDto, error) {
	video, err := domain.NewVideo(data)
	if err != nil {
		return nil, err
	}

	if err := v.videoRepository.Create(*video); err != nil {
		return nil, err
	}

	return video.MapTo(), nil
}

func (v VideosUseCase) Update(id *vo.UniqueEntityID, data domain.VideoInput) (*domain.VideoDto, error) {
	video, err := v.videoRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(errors.ErrFindByIdVideo)
	}

	if err := video.Fill(data); err != nil {
		return nil, err
	}

	if err := v.videoRepository.Update(*video); err != nil {
		return nil, err
	}

	return video.MapTo(), nil
}

func (v VideosUseCase) Delete(id *vo.UniqueEntityID) error {
	if err := v.videoRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
