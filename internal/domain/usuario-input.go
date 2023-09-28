package domain

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UsuarioInput struct {
	Name     *string `json:"name" validate:"required"`
	Email    *string `json:"email" validate:"required,email"`
	Password *string `json:"password" validate:"required,min=6"`
}

func (u *UsuarioInput) prepare() {
	if u.Name != nil {
		*u.Name = strings.ToLower(strings.TrimSpace(*u.Name))
	}

	if u.Email != nil {
		*u.Email = strings.ToLower(strings.TrimSpace(*u.Email))
	}
}

func (u *UsuarioInput) validate() error {
	u.prepare()
	if err := validate.Struct(u); err != nil {
		return err
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
