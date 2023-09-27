package categorias

import (
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"github.com/lncitador/alura-flix-backend/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCategoriaHandlers_update(t *testing.T) {
	sut := setupSut()
	id := sut.ID
	url := sut.URL + "/"

	sut.r.PUT(url+":id", sut.update)

	t.Run("should update a categoria", func(t *testing.T) {
		updateName := helpers.Body{
			"nome": "updated name",
		}

		req, _ := http.NewRequest("PUT", url+id, helpers.MappedJson(updateName))
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should not update a categoria with invalid id", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", url+"invalid", nil)
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("should not update a categoria with id not found", func(t *testing.T) {
		newId, _ := vo.NewUniqueEntityID(nil)

		updateName := helpers.Body{
			"nome": "updated name",
		}

		req, _ := http.NewRequest("PUT", url+newId.ToString(), updateName.MappedJson())
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("should not update a categoria with invalid data", func(t *testing.T) {
		updateName := helpers.Body{
			"name": "teste",
		}

		req, _ := http.NewRequest("PUT", url+id, updateName.MappedJson())
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
}
