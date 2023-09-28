package usecases

import (
	"github.com/lncitador/alura-flix-backend/internal/application/repositories"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

type UsuariosUseCase struct {
	repositories.UsuarioRepositoryContract
}

func NewUsuariosUseCase(repository repositories.UsuarioRepositoryContract) *UsuariosUseCase {
	return &UsuariosUseCase{repository}
}

func (u *UsuariosUseCase) FindAll() (*[]domain.UsuarioDto, error) {
	usuarios, err := u.UsuarioRepositoryContract.FindAll(nil)
	if err != nil {
		return nil, err
	}

	return domain.UsuariosToDto(usuarios), nil
}

func (u *UsuariosUseCase) FindById(id *vo.UniqueEntityID) (*domain.UsuarioDto, error) {
	usuario, err := u.UsuarioRepositoryContract.FindById(id)
	if err != nil {
		return nil, err
	}

	return usuario.MapTo(), nil
}

func (u *UsuariosUseCase) Create(data domain.UsuarioInput) (*domain.UsuarioDto, error) {
	usuario, err := domain.NewUsuario(data)
	if err != nil {
		return nil, err
	}

	if err := u.UsuarioRepositoryContract.Create(*usuario); err != nil {
		return nil, err
	}

	return usuario.MapTo(), nil
}

func (u *UsuariosUseCase) Update(id *vo.UniqueEntityID, data domain.UsuarioInput) (*domain.UsuarioDto, error) {
	usuario, err := u.UsuarioRepositoryContract.FindById(id)
	if err != nil {
		return nil, err
	}

	if err := usuario.Fill(data); err != nil {
		return nil, err
	}

	if err := u.UsuarioRepositoryContract.Update(*usuario); err != nil {
		return nil, err
	}

	return usuario.MapTo(), nil
}

func (u *UsuariosUseCase) Delete(id *vo.UniqueEntityID) error {
	if err := u.UsuarioRepositoryContract.Delete(id); err != nil {
		return err
	}

	return nil
}
