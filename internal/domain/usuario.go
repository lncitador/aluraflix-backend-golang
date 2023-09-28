package domain

type Usuario struct {
	Base
	Nome     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
}

func NewUsuario(input UsuarioInput) (*Usuario, error) {
	usuario := Usuario{}
	usuario.prepare()

	if err := input.validate(); err != nil {
		return nil, err
	}

	usuario.Nome = *input.Nome
	usuario.Email = *input.Email
	usuario.Password = *input.Password

	return &usuario, nil
}

func (u *Usuario) MapTo() *UsuarioDto {
	return &UsuarioDto{
		ID:        u.ID.ToString(),
		Nome:      u.Nome,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *Usuario) Fill(input UsuarioInput) error {
	if err := input.validate(); err != nil {
		return err
	}

	if input.Nome != nil {
		u.Nome = *input.Nome
	}

	if input.Email != nil {
		u.Email = *input.Email
	}

	if input.Password != nil {
		u.Password = *input.Password
	}

	return nil
}
