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

func TestVideoQuery_Page(t *testing.T) {
	t.Run("should return nil if page not set", func(t *testing.T) {
		q := &VideoQuery{}
		assert.Nil(t, q.Page())
	})

	t.Run("should return page value from query", func(t *testing.T) {
		value := 2
		q := &VideoQuery{page: &value}
		assert.Equal(t, &value, q.Page())
	})
}

func TestVideoQuery_SetPage(t *testing.T) {
	t.Run("should set page value from query", func(t *testing.T) {
		q := &VideoQuery{}
		_ = q.SetPage("2")

		assert.Equal(t, 2, *q.page)
	})

	t.Run("should return error if page is not a number", func(t *testing.T) {
		q := &VideoQuery{}
		err := q.SetPage("a")

		assert.EqualError(t, err, "page must be a number")
	})

	t.Run("should return error if page is less than 1", func(t *testing.T) {
		q := &VideoQuery{}
		err := q.SetPage("0")

		assert.EqualError(t, err, "page must be greater than 0")
	})
}
