package categorias

import (
	"encoding/json"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCategoriaHandlers_index(t *testing.T) {
	sut := setupSut()
	sut.r.GET(sut.URL, sut.index)

	t.Run("should return all categorias", func(t *testing.T) {
		req, _ := http.NewRequest("GET", sut.URL, nil)
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var body []domain.CategoriaDto

		if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
			t.Fatalf("could not unmarshal body: %v", err)
		}

		assert.Equal(t, 1, len(body))
		assert.Equal(t, sut.ID, body[0].ID)
	})
}
