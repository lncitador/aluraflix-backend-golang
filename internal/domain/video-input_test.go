package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVideoInput_prepare(t *testing.T) {
	t.Run("should prepare the video input", func(t *testing.T) {
		title := "  My Title  "
		description := "  My Description  "
		url := "  https://my-url.com  "
		categoryID := "  123e4567-e89b-12d3-a456-426614174000  "

		i := &VideoInput{
			Title:       &title,
			Description: &description,
			URL:         &url,
			CategoryID:  &categoryID,
		}

		i.prepare()

		assert.Equal(t, "my title", *i.Title)
		assert.Equal(t, "my description", *i.Description)
		assert.Equal(t, "https://my-url.com", *i.URL)
		assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", *i.CategoryID)
	})
}

func TestVideoInput_validate(t *testing.T) {
	t.Run("should return nil if the video input is valid", func(t *testing.T) {
		title := "  My Title  "
		description := "  My Description  "
		url := "  https://my-url.com  "
		categoryID := "  79c741ed-54fb-4501-8b65-8e9bec766709  "
		usuarioID := "  79c741ed-54fb-4501-8b65-8e9bec766709     "

		i := &VideoInput{
			Title:       &title,
			Description: &description,
			URL:         &url,
			CategoryID:  &categoryID,
			UsuarioID:   &usuarioID,
		}

		err := i.validate()
		assert.Nil(t, err)
	})

	t.Run("should return an error if the video input is invalid", func(t *testing.T) {
		i := &VideoInput{}

		err := i.validate()
		assert.NotNil(t, err)
	})
}
