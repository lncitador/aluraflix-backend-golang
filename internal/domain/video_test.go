package domain

import (
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewVideo(t *testing.T) {
	id, _ := vo.NewUniqueEntityID(nil)
	title := "New video"
	description := "New video description"
	url := "https://www.youtube.com/watch?v=123456789"
	categoryID := id.ToString()

	t.Run("should create a new video", func(t *testing.T) {
		video, err := NewVideo(VideoInput{&title, &description, &url, &categoryID})
		assert.Nil(t, err)
		assert.NotNil(t, video)
	})

	t.Run("should lowercase the title", func(t *testing.T) {
		title := "New video"
		video, _ := NewVideo(VideoInput{&title, &description, &url, &categoryID})
		assert.Equal(t, "new video", video.Title)
	})

	t.Run("should not create a new video with invalid title", func(t *testing.T) {
		video, err := NewVideo(VideoInput{nil, &description, &url, &categoryID})
		assert.NotNil(t, err)
		assert.Nil(t, video)

		invalidTitle := "Bao"

		video, err = NewVideo(VideoInput{&invalidTitle, &description, &url, &categoryID})
		assert.NotNil(t, err)
		assert.Nil(t, video)
	})

	t.Run("should not create a new video with invalid description", func(t *testing.T) {
		video, err := NewVideo(VideoInput{&title, nil, &url, &categoryID})
		assert.NotNil(t, err)
		assert.Nil(t, video)

		invalidDescription := "Bao"

		video, err = NewVideo(VideoInput{&title, &invalidDescription, &url, &categoryID})
		assert.NotNil(t, err)
		assert.Nil(t, video)
	})

	t.Run("should not create a new video with invalid url", func(t *testing.T) {
		video, err := NewVideo(VideoInput{&title, &description, nil, &categoryID})
		assert.NotNil(t, err)
		assert.Nil(t, video)

		invalidUrl := "www.youtube"

		video, err = NewVideo(VideoInput{&title, &description, &invalidUrl, &categoryID})
		assert.NotNil(t, err)
		assert.Nil(t, video)
	})

	t.Run("should not create a new video with invalid category id", func(t *testing.T) {
		video, err := NewVideo(VideoInput{&title, &description, &url, nil})
		assert.NotNil(t, err)
		assert.Nil(t, video)

		invalidCategoryID := "123"

		video, err = NewVideo(VideoInput{&title, &description, &url, &invalidCategoryID})
		assert.NotNil(t, err)
		assert.Nil(t, video)
	})

	t.Run("should not create a new video with invalid data", func(t *testing.T) {
		video, err := NewVideo(VideoInput{})
		assert.NotNil(t, err)
		assert.Nil(t, video)
	})
}
