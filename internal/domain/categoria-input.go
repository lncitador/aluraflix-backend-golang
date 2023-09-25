package domain

import "strings"

type CategoriaInput struct {
	Name  *string `json:"name" validate:"required,min=8,max=255"`
	Color *string `json:"color" validate:"required,min=7,max=7,hexcolor"`
}

func (i *CategoriaInput) prepare() {
	if i.Name != nil {
		*i.Name = strings.ToLower(strings.TrimSpace(*i.Name))
	}
}

func (i *CategoriaInput) validate() error {
	if err := validate.Struct(i); err != nil {
		return err
	}
	i.prepare()

	return nil
}
