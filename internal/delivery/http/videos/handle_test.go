package videos

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories/inmemory"
	"github.com/lncitador/alura-flix-backend/internal/application/usecases"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	"log"
)

type Constants struct {
	Truth   domain.VideoInput
	Falsy   domain.VideoInput
	Invalid domain.VideoInput
}

type Sut struct {
	*VideoHandlers
	URL       string
	ID        string
	r         *gin.Engine
	constants Constants
}

func setupSut() *Sut {
	r := gin.Default()

	relation := inmemory.NewCategoriaRepository()

	name := "Categoria Teste"
	color := "#000000"

	var usuarioId string

	{
		name := "John Doe"
		email := "doejoe@test.com"
		password := "123456"

		usuario, _ := domain.NewUsuario(domain.UsuarioInput{
			Name:     &name,
			Email:    &email,
			Password: &password,
		})

		usuarioId = usuario.ID.ToString()
	}

	categoria, _ := domain.NewCategoria(domain.CategoriaInput{
		Name:      &name,
		Color:     &color,
		UsuarioID: &usuarioId,
	})

	if err := relation.Create(*categoria); err != nil {
		log.Fatalf("could not create categoria: %v", err.Error())
	}

	repo := inmemory.NewVideoRepository(relation)

	useCase := usecases.NewVideosUseCase(repo)

	router := r.Group("/api")

	h := &VideoHandlers{
		router:  router,
		useCase: useCase,
	}

	truthTitle := "Video Teste"
	truthDescription := "Video Teste"
	truthURL := "https://www.youtube.com/watch?v=1"
	truthCategoryID := categoria.ID.ToString()
	invalidTitle := "Video"
	invalidDescription := "Video"
	invalidURL := "https://www.youtube.com/watch?v=1"
	invalidCategoryID := "1"
	invalidUsuarioID := "1"

	constants := Constants{
		Truth: domain.VideoInput{
			Title:       &truthTitle,
			Description: &truthDescription,
			URL:         &truthURL,
			CategoryID:  &truthCategoryID,
			UsuarioID:   &usuarioId,
		},
		Falsy: domain.VideoInput{
			Title:       nil,
			Description: nil,
			URL:         nil,
			CategoryID:  nil,
			UsuarioID:   nil,
		},
		Invalid: domain.VideoInput{
			Title:       &invalidTitle,
			Description: &invalidDescription,
			URL:         &invalidURL,
			CategoryID:  &invalidCategoryID,
			UsuarioID:   &invalidUsuarioID,
		},
	}

	video, err := useCase.Create(constants.Truth)
	if err != nil {
		return nil
	}

	r.Use(func(c *gin.Context) {
		c.Set("user", &domain.UsuarioDto{
			ID: usuarioId,
		})
	})

	return &Sut{
		VideoHandlers: h,
		URL:           "/api/v1/videos",
		ID:            video.ID,
		r:             r,
		constants:     constants,
	}
}
