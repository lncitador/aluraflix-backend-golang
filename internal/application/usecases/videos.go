package usecases

import (
	"fmt"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

type VideoRepositoryContract repositories.RepositoryContract[domain.Video, domain.VideoQuery]

type VideosUseCase struct {
	VideoRepositoryContract
}

func NewVideosUseCase(contract VideoRepositoryContract) *VideosUseCase {
	return &VideosUseCase{contract}
}

func (v VideosUseCase) FindAll(query domain.VideoQuery) (*[]domain.VideoDto, error) {
	videos, err := v.VideoRepositoryContract.FindAll(query)
	if err != nil {
		return nil, err
	}

	return domain.VideosToDto(videos), nil
}

func (v VideosUseCase) FindById(id *vo.UniqueEntityID) (*domain.VideoDto, error) {
	video, err := v.VideoRepositoryContract.FindById(id)
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

	if err := v.VideoRepositoryContract.Create(*video); err != nil {
		return nil, err
	}

	return video.MapTo(), nil
}

func (v VideosUseCase) Update(id *vo.UniqueEntityID, data domain.VideoInput) (*domain.VideoDto, error) {
	video, err := v.VideoRepositoryContract.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(errors.ErrFindByIdVideo)
	}

	if err := video.Fill(data); err != nil {
		return nil, err
	}

	if err := v.VideoRepositoryContract.Update(*video); err != nil {
		return nil, err
	}

	return video.MapTo(), nil
}

func (v VideosUseCase) Delete(id *vo.UniqueEntityID) error {
	if err := v.VideoRepositoryContract.Delete(id); err != nil {
		return err
	}

	return nil
}
