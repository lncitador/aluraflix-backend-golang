package videos

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/config"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories"
	"github.com/lncitador/alura-flix-backend/internal/application/usecases"
)

type VideoHandlers struct {
	router  *gin.Engine
	useCase *usecases.VideosUseCase
}

func NewVideoHandlers(config *config.Config) *VideoHandlers {
	db := config.GetDb()

	repo := repositories.NewVideoRepository(db)
	useCase := usecases.NewVideosUseCase(repo)

	return &VideoHandlers{
		router:  config.GetRouter(),
		useCase: useCase,
	}
}
