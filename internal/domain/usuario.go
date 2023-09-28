package domain

import "github.com/go-playground/validator/v10"

type Usuario struct {
	Base
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
}

func NewUsuario(input UsuarioInput) (*Usuario, error) {
	usuario := Usuario{}
	usuario.prepare()

	if err := input.validate(); err != nil {
		return nil, err
	}

	usuario.Name = *input.Name
	usuario.Email = *input.Email
	usuario.Password = *input.Password

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

func (u *Usuario) Fill(input UsuarioInput) error {
	err := validate.Struct(input)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				continue
			}
			return err
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
