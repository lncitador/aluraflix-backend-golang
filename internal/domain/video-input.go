package domain

import "strings"

type VideoInput struct {
	Title       *string `json:"title" validate:"required,min=8,max=255"`
	Description *string `json:"description" validate:"required,min=8,max=255"`
	URL         *string `json:"url" validate:"required,url"`
	CategoryID  *string `json:"categoryID" validate:"required,uuid4"`
	UsuarioID   *string `json:"usuarioID" validate:"required,uuid4"`
}

func (i *VideoInput) prepare() {
	if i.Title != nil {
		*i.Title = strings.ToLower(strings.TrimSpace(*i.Title))
	}

	if i.Description != nil {
		*i.Description = strings.ToLower(strings.TrimSpace(*i.Description))
	}

	if i.URL != nil {
		*i.URL = strings.ToLower(strings.TrimSpace(*i.URL))
	}

	if i.CategoryID != nil {
		*i.CategoryID = strings.ToLower(strings.TrimSpace(*i.CategoryID))
	}

	if i.UsuarioID != nil {
		*i.UsuarioID = strings.ToLower(strings.TrimSpace(*i.UsuarioID))
	}
}

func (i *VideoInput) validate() error {
	i.prepare()

	if err := validate.Struct(i); err != nil {
		return err
	}

	return nil
}
