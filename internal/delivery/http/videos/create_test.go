package videos

import (
	"github.com/lncitador/alura-flix-backend/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVideoHandlers_create(t *testing.T) {
	sut := setupSut()
	sut.r.POST(sut.URL, sut.create)

	data := sut.constants

	t.Run("should create a video", func(t *testing.T) {
		req, _ := http.NewRequest("POST", sut.URL, helpers.MappedJson(data.Truth))
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should not create a video with invalid data", func(t *testing.T) {
		req, _ := http.NewRequest("POST", sut.URL, helpers.MappedJson(data.Invalid))
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("should not create a video with invalid category", func(t *testing.T) {
		body := helpers.Body{
			"title":       data.Truth.Title,
			"description": data.Truth.Description,
			"url":         data.Truth.URL,
			"categoryId":  "1",
		}

		req, _ := http.NewRequest("POST", sut.URL, body.MappedJson())
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
}
