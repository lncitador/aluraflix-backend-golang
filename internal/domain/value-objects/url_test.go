package value_objects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewURL(t *testing.T) {
	t.Run("should create a new url", func(t *testing.T) {
		url, err := NewURL("https://www.youtube.com/watch?v=123456789")
		assert.Nil(t, err)
		assert.NotNil(t, url)
	})

	t.Run("should not create a new url with invalid url", func(t *testing.T) {
		url, err := NewURL("www.youtube")
		assert.NotNil(t, err)
		assert.Nil(t, url)
	})
}
