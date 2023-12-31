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

		assert.Equal(t, "page must be a number", err.GetDescription())
	})

	t.Run("should return error if page is less than 1", func(t *testing.T) {
		q := &VideoQuery{}
		err := q.SetPage("0")

		assert.Equal(t, "page must be greater than 0", err.GetDescription())
	})
}

func TestVideoQuery_Limit(t *testing.T) {
	t.Run("should return default limit if limit not set", func(t *testing.T) {
		q := &VideoQuery{}
		assert.Equal(t, 10, *q.Limit())
	})

	t.Run("should return limit value from query", func(t *testing.T) {
		value := 20
		q := &VideoQuery{limit: &value}
		assert.Equal(t, &value, q.Limit())
	})
}

func TestVideoQuery_SetLimit(t *testing.T) {
	t.Run("should set limit value from query", func(t *testing.T) {
		q := &VideoQuery{}
		_ = q.SetLimit("20")

		assert.Equal(t, 20, *q.limit)
	})

	t.Run("should return error if limit is not a number", func(t *testing.T) {
		q := &VideoQuery{}
		err := q.SetLimit("a")

		assert.Equal(t, "limit must be a number", err.GetDescription())
	})

	t.Run("should return error if limit is less than 10", func(t *testing.T) {
		q := &VideoQuery{}
		err := q.SetLimit("9")

		assert.Equal(t, "limit must be greater than 10", err.GetDescription())
	})

	t.Run("should return error if limit is greater than 100", func(t *testing.T) {
		q := &VideoQuery{}
		err := q.SetLimit("101")

		assert.Equal(t, "limit must be less than 100", err.GetDescription())
	})
}

func TestVideoQuery_SetTotal(t *testing.T) {
	t.Run("should set total value from query", func(t *testing.T) {
		q := &VideoQuery{}
		q.SetTotal(10)

		assert.Equal(t, int64(10), *q.total)
	})
}

func TestVideoQuery_Total(t *testing.T) {
	t.Run("should return total value from query", func(t *testing.T) {
		value := int64(10)
		q := &VideoQuery{total: &value}
		assert.Equal(t, &value, q.Total())
	})
}
