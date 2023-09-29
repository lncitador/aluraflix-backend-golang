package usecases

import (
	"github.com/lncitador/alura-flix-backend/internal/application/repositories"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	. "github.com/lncitador/alura-flix-backend/pkg/errors"
)

type UsuariosUseCase struct {
	repositories.UsuarioRepositoryContract
}

func NewUsuariosUseCase(repository repositories.UsuarioRepositoryContract) *UsuariosUseCase {
	return &UsuariosUseCase{repository}
}

func (u *UsuariosUseCase) FindAll() (*[]domain.UsuarioDto, Error) {
	usuarios, err := u.UsuarioRepositoryContract.FindAll(nil)
	if err != nil {
		return nil, err
	}

	return domain.UsuariosToDto(usuarios), nil
}

func (u *UsuariosUseCase) FindById(id *vo.UniqueEntityID) (*domain.UsuarioDto, Error) {
	usuario, err := u.UsuarioRepositoryContract.FindById(id)
	if err != nil {
		return nil, err
	}

	return usuario.MapTo(), nil
}

func (u *UsuariosUseCase) Create(data domain.UsuarioInput) (*domain.UsuarioDto, Error) {
	usuario, err := domain.NewUsuario(data)
	if err != nil {
		return nil, err
	}

	if err := u.UsuarioRepositoryContract.Create(*usuario); err != nil {
		return nil, err
	}

	return usuario.MapTo(), nil
}

func (u *UsuariosUseCase) Update(id *vo.UniqueEntityID, data domain.UsuarioInput) (*domain.UsuarioDto, Error) {
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

func (u *UsuariosUseCase) Delete(id *vo.UniqueEntityID) Error {
	if err := u.UsuarioRepositoryContract.Delete(id); err != nil {
		return err
	}

	return nil
}

func (u *UsuariosUseCase) Signin(email string, password string) (*string, Error) {
	credentials, err := domain.NewCredential(email, password)
	if err != nil {
		return nil, err
	}

	usuario, err := u.UsuarioRepositoryContract.FindByEmail(credentials.Email)
	if err != nil {
		return nil, err
	}

	if err := usuario.ComparePassword(credentials.Password); err != nil {
		return nil, err
	}

	token, err := usuario.GenerateToken()
	if err != nil {
		return nil, err
	}

	return token, nil
}
