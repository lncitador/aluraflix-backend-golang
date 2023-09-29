package domain

import "github.com/lncitador/alura-flix-backend/pkg/errors"

type Credentials struct {
	Email    *string `json:"email" validate:"required,email"`
	Password *string `json:"password" validate:"required,min=6"`
}

func NewCredential(email string, password string) (*Credentials, errors.Error) {
	credential := &Credentials{
		Email:    &email,
		Password: &password,
	}

	if err := credential.validate(); err != nil {
		return nil, errors.NewErrorByValidation(err)
	}

	return credential, nil
}

func (c *Credentials) validate() error {
	return validate.Struct(c)
}
