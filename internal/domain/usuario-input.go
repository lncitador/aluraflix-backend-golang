package domain

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UsuarioInput struct {
	Nome     *string `json:"nome" validate:"required"`
	Email    *string `json:"email" validate:"required,email"`
	Password *string `json:"password" validate:"required,min=6"`
}

func (u *UsuarioInput) prepare() error {
	if u.Nome != nil {
		*u.Nome = strings.ToLower(strings.TrimSpace(*u.Nome))
	}

	if u.Email != nil {
		*u.Email = strings.ToLower(strings.TrimSpace(*u.Email))
	}

	if u.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*u.Password), bcrypt.DefaultCost)

		if err != nil {
			return err
		}

		*u.Password = string(hash)
	}

	return nil
}

func (u *UsuarioInput) validate() error {
	if err := u.prepare(); err != nil {
		return err
	}

	return validate.Struct(u)
}
