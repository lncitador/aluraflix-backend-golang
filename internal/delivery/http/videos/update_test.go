package videos

import (
	"github.com/lncitador/alura-flix-backend/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVideoHandlers_update(t *testing.T) {
	sut := setupSut()
	url := sut.URL + "/"
	sut.r.PUT(url+":id", sut.update)

	id := sut.ID

	t.Run("should update a video", func(t *testing.T) {
		body := helpers.Body{
			"title": "new title",
		}

		req, _ := http.NewRequest("PUT", url+id, body.MappedJson())
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should not update a video with invalid id", func(t *testing.T) {
		body := helpers.Body{
			"title": "new title",
		}

		req, _ := http.NewRequest("PUT", url+"invalid", body.MappedJson())
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("should not update a video with id not found", func(t *testing.T) {
		body := helpers.Body{
			"title": "new title",
		}

		req, _ := http.NewRequest("PUT", url+"8f3c047e-259a-48e4-b357-aa04f00e9d58", body.MappedJson())
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("should not update a video with nil body", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", url+id, nil)
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should not update a video with invalid body", func(t *testing.T) {
		body := helpers.Body{
			"title": "",
		}

		req, _ := http.NewRequest("PUT", url+id, body.MappedJson())
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
}
