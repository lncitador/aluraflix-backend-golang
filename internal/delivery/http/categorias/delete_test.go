package categorias

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCategoriaHandlers_delete(t *testing.T) {
	sut := setupSut()
	id := sut.ID
	url := sut.URL + "/"

	sut.r.DELETE(url+":id", sut.delete)

	t.Run("should delete a categoria", func(t *testing.T) {

		req, _ := http.NewRequest("DELETE", url+id, nil)
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("should not delete a categoria with invalid id", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", url+"invalid", nil)
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
