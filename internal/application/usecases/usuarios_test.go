package usecases

import (
	"github.com/lncitador/alura-flix-backend/internal/application/repositories/inmemory"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"github.com/lncitador/alura-flix-backend/pkg/validations"
	"github.com/stretchr/testify/assert"
	"testing"
)

type UsuarioSut struct {
	repo      *inmemory.UsuariosRepository
	useCase   *UsuariosUseCase
	constants domain.UsuarioInput
}

func setupSut() *UsuarioSut {
	repo := inmemory.NewUsuariosRepository()

	useCase := NewUsuariosUseCase(repo)

	name := "John Doe"
	email := "doejoe@test.com"
	password := "123456"

	return &UsuarioSut{
		repo:    repo,
		useCase: useCase,
		constants: domain.UsuarioInput{
			Name:     &name,
			Email:    &email,
			Password: &password,
		},
	}
}

func TestUsuariosUseCase_Create(t *testing.T) {
	sut := setupSut()
	data := sut.constants

	t.Run("should create a user", func(t *testing.T) {
		got, err := sut.useCase.Create(data)

		assert.Nil(t, err)
		assert.NotNil(t, got)
	})

	t.Run("should not create a user with invalid email", func(t *testing.T) {
		invalidEmail := "doejoe@test"

		got, err := sut.useCase.Create(domain.UsuarioInput{
			Name:     data.Name,
			Email:    &invalidEmail,
			Password: data.Password,
		})

		assert.NotNil(t, err)
		assert.Nil(t, got)

		isErrValidation, _ := validations.ErrosAsValidationByField(err, "Email")

		assert.True(t, *isErrValidation)
	})

	t.Run("should not create a user with invalid password", func(t *testing.T) {
		invalidPassword := "123"

		got, err := sut.useCase.Create(domain.UsuarioInput{
			Name:     data.Name,
			Email:    data.Email,
			Password: &invalidPassword,
		})

		assert.NotNil(t, err)
		assert.Nil(t, got)

		isErrValidation, _ := validations.ErrosAsValidationByField(err, "Password")

		assert.True(t, *isErrValidation)
	})
}

func TestUsuariosUseCase_Delete(t *testing.T) {
	sut := setupSut()

	t.Run("should delete a user", func(t *testing.T) {
		got, err := sut.useCase.Create(sut.constants)

		id, err := vo.NewUniqueEntityID(&got.ID)
		assert.Nil(t, err)
		assert.NotNil(t, got)

		err = sut.useCase.Delete(id)
		assert.Nil(t, err)
	})
}

func TestUsuariosUseCase_FindAll(t *testing.T) {
	sut := setupSut()
	data := sut.constants

	t.Run("should find all users", func(t *testing.T) {
		got, err := sut.useCase.FindAll()

		assert.Nil(t, err)
		assert.NotNil(t, got)

		assert.Equal(t, 0, len(*got))

		_, err = sut.useCase.Create(data)
		assert.Nil(t, err)

		got, err = sut.useCase.FindAll()

		assert.Nil(t, err)
		assert.NotNil(t, got)

		assert.Equal(t, 1, len(*got))
	})
}

func TestUsuariosUseCase_FindById(t *testing.T) {
	sut := setupSut()
	data := sut.constants

	t.Run("should find a user by id", func(t *testing.T) {
		got, err := sut.useCase.Create(data)

		id, err := vo.NewUniqueEntityID(&got.ID)
		assert.Nil(t, err)
		assert.NotNil(t, got)

		got, err = sut.useCase.FindById(id)

		assert.Nil(t, err)
		assert.NotNil(t, got)
	})

	t.Run("should not find a user by not found id", func(t *testing.T) {
		id, err := vo.NewUniqueEntityID(nil)
		assert.Nil(t, err)

		got, err := sut.useCase.FindById(id)

		assert.Nil(t, err)
		assert.Nil(t, got)
	})

	t.Run("should not find a user by invalid id", func(t *testing.T) {
		got, err := sut.useCase.FindById(nil)

		assert.Nil(t, err)
		assert.Nil(t, got)
	})
}

func TestUsuariosUseCase_Update(t *testing.T) {
	sut := setupSut()
	data := sut.constants

	t.Run("should update a user", func(t *testing.T) {
		got, err := sut.useCase.Create(data)

		id, err := vo.NewUniqueEntityID(&got.ID)
		assert.Nil(t, err)
		assert.NotNil(t, got)

		newName := "John Doe 2"

		data := domain.UsuarioInput{
			Name: &newName,
		}

		got, err = sut.useCase.Update(id, data)

		assert.Nil(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, *data.Name, got.Name)
	})

	t.Run("should not update a user with invalid email", func(t *testing.T) {
		got, err := sut.useCase.Create(data)

		id, err := vo.NewUniqueEntityID(&got.ID)
		assert.Nil(t, err)
		assert.NotNil(t, got)

		invalidEmail := "doejoe@test"

		data := domain.UsuarioInput{
			Email: &invalidEmail,
		}

		got, err = sut.useCase.Update(id, data)

		assert.NotNil(t, err)
		assert.Nil(t, got)

		isErrValidation, _ := validations.ErrosAsValidationByField(err, "Email")

		assert.True(t, *isErrValidation)
	})

	t.Run("should not update a user with invalid password", func(t *testing.T) {
		got, err := sut.useCase.Create(data)

		id, err := vo.NewUniqueEntityID(&got.ID)
		assert.Nil(t, err)
		assert.NotNil(t, got)

		invalidPassword := "123"

		data := domain.UsuarioInput{
			Password: &invalidPassword,
		}

		got, err = sut.useCase.Update(id, data)

		assert.NotNil(t, err)
		assert.Nil(t, got)

		isErrValidation, _ := validations.ErrosAsValidationByField(err, "Password")

		assert.True(t, *isErrValidation)
	})

	t.Run("should not update a user with invalid id", func(t *testing.T) {
		got, err := sut.useCase.Create(data)

		id, err := vo.NewUniqueEntityID(nil)
		assert.Nil(t, err)
		assert.NotNil(t, got)

		newName := "John Doe 2"

		data := domain.UsuarioInput{
			Name: &newName,
		}

		got, err = sut.useCase.Update(id, data)

		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
}
