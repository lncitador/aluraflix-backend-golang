package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUsuario(t *testing.T) {
	name := "John Doe"
	email := "doejoe@test.com"
	password := "123456"

	t.Run("should create a new user", func(t *testing.T) {
		got, err := NewUsuario(UsuarioInput{
			Name:     &name,
			Email:    &email,
			Password: &password,
		})

		assert.Nil(t, err)
		assert.NotNil(t, got)
	})

	t.Run("should not create a new user with invalid email", func(t *testing.T) {
		invalidMail := "doejoe@test"

		got, err := NewUsuario(UsuarioInput{
			Name:     &name,
			Email:    &invalidMail,
			Password: &password,
		})

		assert.NotNil(t, err)
		assert.Nil(t, got)
	})

	t.Run("should not create a new user with short password", func(t *testing.T) {
		shortPassword := "123"

		got, err := NewUsuario(UsuarioInput{
			Name:     &name,
			Email:    &email,
			Password: &shortPassword,
		})

		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
}

func TestUsuario_ComparePassword(t *testing.T) {
	name := "John Doe"
	email := "doejoe@test.com"
	password := "123456"

	t.Run("should compare password", func(t *testing.T) {
		u, _ := NewUsuario(UsuarioInput{
			Name:     &name,
			Email:    &email,
			Password: &password,
		})

		err := u.ComparePassword(&password)

		assert.Nil(t, err)
	})

	t.Run("should not compare password", func(t *testing.T) {
		invalidPwd := "123"
		u, _ := NewUsuario(UsuarioInput{
			Name:     &name,
			Email:    &email,
			Password: &password,
		})

		err := u.ComparePassword(&invalidPwd)

		assert.NotNil(t, err)
	})
}

func TestUsuario_GenerateToken(t *testing.T) {
	name := "John Doe"
	email := "doejoe@test.com"
	password := "123456"

	t.Run("should generate token", func(t *testing.T) {
		u, _ := NewUsuario(UsuarioInput{
			Name:     &name,
			Email:    &email,
			Password: &password,
		})

		token, err := u.GenerateToken()

		assert.Nil(t, err)
		assert.NotNil(t, token)
	})
}
