package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVideoQuery_Search(t *testing.T) {
	t.Run("should return nil if search is nil", func(t *testing.T) {
		q := &VideoQuery{}
		assert.Nil(t, q.Search())
	})

	t.Run("should return search value from query", func(t *testing.T) {
		value := "%search%"
		q := &VideoQuery{search: &value}
		assert.Equal(t, &value, q.Search())
	})
}

func TestVideoQuery_SetSearch(t *testing.T) {
	t.Run("should set search value from query", func(t *testing.T) {
		q := &VideoQuery{}
		q.SetSearch("search")

		assert.Equal(t, "%search%", *q.search)
	})
}
