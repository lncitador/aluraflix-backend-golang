package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories/inmemory"
	"github.com/lncitador/alura-flix-backend/internal/application/usecases"
	"github.com/lncitador/alura-flix-backend/internal/domain"
)

type Constants struct {
	Truth   domain.Credentials
	Falsy   domain.Credentials
	Invalid domain.Credentials
}

type Sut struct {
	*AuthHandlers
	URL       string
	ID        string
	r         *gin.Engine
	constants Constants
}

func setupSut() *Sut {
	r := gin.Default()

	repo := inmemory.NewUsuariosRepository()
	useCase := usecases.NewUsuariosUseCase(repo)

	router := r.Group("/api")

	h := &AuthHandlers{
		router:  router,
		useCase: useCase,
	}

	truthEmail := "doejoe@test.com"
	truthPassword := "123456"
	invalidEmail := "doejoe@test"
	invalidPassword := "123"

	constants := Constants{
		Truth: domain.Credentials{
			Email:    &truthEmail,
			Password: &truthPassword,
		},
		Falsy: domain.Credentials{
			Email:    nil,
			Password: nil,
		},
		Invalid: domain.Credentials{
			Email:    &invalidEmail,
			Password: &invalidPassword,
		},
	}

	name := "Joe Doe"

	usuario, _ := useCase.Create(domain.UsuarioInput{
		Name:     &name,
		Email:    &truthEmail,
		Password: &truthPassword,
	})

	return &Sut{
		AuthHandlers: h,
		URL:          "/api/v1/auth",
		ID:           usuario.ID,
		r:            r,
		constants:    constants,
	}
}
