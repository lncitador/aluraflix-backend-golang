package authorization

import (
	"github.com/lncitador/alura-flix-backend/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHandlers_signup(t *testing.T) {
	sut := setupSut()
	data := sut.constants

	sut.r.POST(sut.URL, sut.signup)

	t.Run("should return 400 if dto is invalid", func(t *testing.T) {
		req, _ := http.NewRequest("POST", sut.URL, helpers.MappedJson(data.Falsy))
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("should return 201 if dto is valid", func(t *testing.T) {
		dto := helpers.Body{
			"name":     "Mary Joe",
			"email":    "jojo@test.com",
			"password": "123456",
		}

		req, _ := http.NewRequest("POST", sut.URL, dto.MappedJson())
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})
}
