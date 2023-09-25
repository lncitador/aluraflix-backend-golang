package domain

import "strings"

type VideoInput struct {
	Title       *string `json:"title" validate:"required,min=8,max=255"`
	Description *string `json:"description" validate:"required,min=8,max=255"`
	URL         *string `json:"url" validate:"required,url"`
	CategoryID  *string `json:"categoryID" validate:"required,uuid4"`
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
}

func (i *VideoInput) validate() error {
	if err := validate.Struct(i); err != nil {
		return err
	}

	i.prepare()

	return nil
}
