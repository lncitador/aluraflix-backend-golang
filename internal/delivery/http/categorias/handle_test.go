package categorias

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories/inmemory"
	"github.com/lncitador/alura-flix-backend/internal/application/usecases"
	"github.com/lncitador/alura-flix-backend/internal/domain"
)

type Constants struct {
	Truth   domain.CategoriaInput
	Falsy   domain.CategoriaInput
	Invalid domain.CategoriaInput
}

type Sut struct {
	*CategoriaHandlers
	URL       string
	ID        string
	r         *gin.Engine
	constants Constants
}

func setupSut() *Sut {
	r := gin.Default()

	repo := inmemory.NewCategoriaRepository()
	useCase := usecases.NewCategoriasUseCase(repo)

	router := r.Group("/api")

	h := &CategoriaHandlers{
		router:  router,
		useCase: useCase,
	}

	truthName := "Categoria Teste"
	truthColor := "#000000"
	invalidName := "Catego"
	invalidColor := "#0000000"

	constants := Constants{
		Truth: domain.CategoriaInput{
			Name:  &truthName,
			Color: &truthColor,
		},
		Falsy: domain.CategoriaInput{
			Name:  nil,
			Color: nil,
		},
		Invalid: domain.CategoriaInput{
			Name:  &invalidName,
			Color: &invalidColor,
		},
	}

	categoria, _ := useCase.Create(constants.Truth)

	return &Sut{
		CategoriaHandlers: h,
		URL:               "/api/v1/categorias",
		ID:                categoria.ID,
		r:                 r,
		constants:         constants,
	}
}
