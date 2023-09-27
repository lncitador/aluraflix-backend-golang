package videos

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVideoHandlers_show(t *testing.T) {
	sut := setupSut()
	url := sut.URL + "/"

	sut.r.GET(url+":id", sut.show)

	id := sut.ID

	t.Run("should return a video", func(t *testing.T) {
		req, _ := http.NewRequest("GET", url+id, nil)
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should not return a video with invalid id", func(t *testing.T) {
		req, _ := http.NewRequest("GET", url+"invalid", nil)
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("should not return a video with id not found", func(t *testing.T) {
		req, _ := http.NewRequest("GET", url+"8f3c047e-259a-48e4-b357-aa04f00e9d58", nil)
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
