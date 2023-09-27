package videos

import (
	"encoding/json"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVideoHandlers_index(t *testing.T) {
	sut := setupSut()
	sut.r.GET(sut.URL, sut.index)

	t.Run("should return a list of videos", func(t *testing.T) {
		req, _ := http.NewRequest("GET", sut.URL, nil)
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return a list of videos with pagination", func(t *testing.T) {
		req, _ := http.NewRequest("GET", sut.URL+"?page=1", nil)
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var body domain.Pagination[domain.VideoDto]

		if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
			t.Fatalf("could not unmarshal body: %v", err)
		}

		assert.Equal(t, 1, len(body.Data))
		assert.Equal(t, int64(1), body.TotalPage)
	})
}
