package authorization

import (
	"github.com/lncitador/alura-flix-backend/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHandlers_signin(t *testing.T) {
	sut := setupSut()
	data := sut.constants

	sut.r.POST(sut.URL, sut.signin)

	t.Run("should signin a user", func(t *testing.T) {
		req, _ := http.NewRequest("POST", sut.URL, helpers.MappedJson(data.Truth))
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should not signin a user with invalid data", func(t *testing.T) {
		req, _ := http.NewRequest("POST", sut.URL, helpers.MappedJson(data.Invalid))
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("should not signin a user with falsy data", func(t *testing.T) {
		req, _ := http.NewRequest("POST", sut.URL, helpers.MappedJson(data.Falsy))
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should not signin a user with invalid email", func(t *testing.T) {
		invalidMail := helpers.Body{
			"email":    "doejoe@test",
			"password": "123456",
		}

		req, _ := http.NewRequest("POST", sut.URL, invalidMail.MappedJson())
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("should not signin a user with invalid password", func(t *testing.T) {
		invalidPassword := helpers.Body{
			"email":    data.Truth.Email,
			"password": "123",
		}

		req, _ := http.NewRequest("POST", sut.URL, invalidPassword.MappedJson())
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("should not signin a user with another user's email", func(t *testing.T) {
		invalidEmail := helpers.Body{
			"email":    "joedoe@test.com",
			"password": data.Truth.Password,
		}

		req, _ := http.NewRequest("POST", sut.URL, invalidEmail.MappedJson())
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("should not signin a user with another user's password", func(t *testing.T) {
		invalidPassword := helpers.Body{
			"email":    data.Truth.Email,
			"password": "654321",
		}

		req, _ := http.NewRequest("POST", sut.URL, invalidPassword.MappedJson())
		w := httptest.NewRecorder()
		sut.r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
