package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lncitador/alura-flix-backend/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type Usuario struct {
	Base
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
}

const TokenMaxAge = time.Hour * 6

func NewUsuario(input UsuarioInput) (*Usuario, errors.Error) {
	usuario := Usuario{}
	usuario.prepare()

	if err := input.validate(); err != nil {
		return nil, errors.NewErrorByValidation(err)
	}

	if input.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)

		if err != nil {
			return nil, errors.NewErrorByInternal(err)
		}

		usuario.Password = string(hash)
	}

	usuario.Name = *input.Name
	usuario.Email = *input.Email

	return &usuario, nil
}

func (u *Usuario) MapTo() *UsuarioDto {
	return &UsuarioDto{
		ID:        u.ID.ToString(),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *Usuario) Fill(input UsuarioInput) errors.Error {
	err := validate.Struct(input)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				continue
			}
			return errors.NewErrorByValidation(err)
		}
	}

	if input.Name != nil {
		u.Name = *input.Name
	}

	if input.Email != nil {
		u.Email = *input.Email
	}

	if input.Password != nil {
		u.Password = *input.Password
	}

	return nil
}

func (u *Usuario) ComparePassword(password *string) errors.Error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(*password)); err != nil {
		return errors.NewErrorByUnauthorized(err)
	}

	return nil
}

func (u *Usuario) GenerateToken() (*string, errors.Error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  u.ID.ToString(),
		"name": u.Name,
		"exp":  TokenMaxAge,
	})

	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, errors.NewErrorByInternal(err)
	}

	return &tokenString, nil
}
