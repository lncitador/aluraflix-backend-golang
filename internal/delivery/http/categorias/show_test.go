package categorias

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCategoriaHandlers_show(t *testing.T) {
	sut := setupSut()
	id := sut.ID
	url := sut.URL + "/"

	sut.r.GET(url+":id", sut.show)

	t.Run("should return a categoria", func(t *testing.T) {

		req, _ := http.NewRequest("GET", url+id, nil)
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should not return a categoria with invalid id", func(t *testing.T) {
		req, _ := http.NewRequest("GET", url+"invalid", nil)
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
}
