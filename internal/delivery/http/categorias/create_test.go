package categorias

import (
	"github.com/lncitador/alura-flix-backend/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCategoriaHandlers_create(t *testing.T) {
	sut := setupSut()
	data := sut.constants

	sut.r.GET(sut.URL, sut.create)

	t.Run("should create a categoria", func(t *testing.T) {
		req, _ := http.NewRequest("GET", sut.URL, helpers.MappedJson(data.Truth))
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should not create a categoria with invalid data", func(t *testing.T) {
		req, _ := http.NewRequest("GET", sut.URL, helpers.MappedJson(data.Invalid))
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
}
