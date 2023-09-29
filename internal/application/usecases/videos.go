package usecases

import (
	"github.com/lncitador/alura-flix-backend/internal/application/repositories"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	. "github.com/lncitador/alura-flix-backend/pkg/errors"
	"net/http"
)

type VideosUseCase struct {
	repositories.VideoRepositoryContract
}

func NewVideosUseCase(contract repositories.VideoRepositoryContract) *VideosUseCase {
	return &VideosUseCase{contract}
}

func (v VideosUseCase) FindAll(query *domain.VideoQuery) (*[]domain.VideoDto, Error) {
	videos, err := v.VideoRepositoryContract.FindAll(query)
	if err != nil {
		return nil, err
	}

	return domain.VideosToDto(videos), nil
}

func (v VideosUseCase) FindById(id *vo.UniqueEntityID, user *vo.UniqueEntityID) (*domain.VideoDto, Error) {
	video, err := v.VideoRepositoryContract.FindById(id)
	if err != nil {
		return nil, err
	}

	if !video.UsuarioID.Equals(user) {
		return nil, NewError(http.StatusUnauthorized, "Unauthorized", "You are not allowed to access this video.")
	}

	return video.MapTo(), nil
}

func (v VideosUseCase) Create(data domain.VideoInput) (*domain.VideoDto, Error) {
	video, err := domain.NewVideo(data)
	if err != nil {
		return nil, err
	}

	if err := v.VideoRepositoryContract.Create(*video); err != nil {
		return nil, err
	}

	return video.MapTo(), nil
}

func (v VideosUseCase) Update(id *vo.UniqueEntityID, data domain.VideoInput) (*domain.VideoDto, Error) {
	video, err := v.VideoRepositoryContract.FindById(id)
	if err != nil {
		return nil, err
	}

	userId, err := vo.NewUniqueEntityID(data.UsuarioID)
	if err != nil {
		return nil, err
	}

	if !video.UsuarioID.Equals(userId) {
		return nil, NewError(http.StatusUnauthorized, "Unauthorized", "You are not allowed to update this video.")
	}

	if err := video.Fill(data); err != nil {
		return nil, err
	}

	if err := v.VideoRepositoryContract.Update(*video); err != nil {
		return nil, err
	}

	return video.MapTo(), nil
}

func (v VideosUseCase) Delete(id *vo.UniqueEntityID, userId *vo.UniqueEntityID) Error {
	video, err := v.VideoRepositoryContract.FindById(id)
	if err != nil {
		return err
	}

	if !video.UsuarioID.Equals(userId) {
		return NewError(http.StatusUnauthorized, "Unauthorized", "You are not allowed to delete this video.")
	}

	if err := v.VideoRepositoryContract.Delete(id); err != nil {
		return err
	}

	return nil
}
