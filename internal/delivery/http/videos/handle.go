package videos

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/config"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories/impl"
	"github.com/lncitador/alura-flix-backend/internal/application/usecases"
)

type VideoHandlers struct {
	router  *gin.Engine
	useCase *usecases.VideosUseCase
}

func NewVideoHandlers(config *config.Config) *VideoHandlers {
	db := config.GetDb()

	repo := impl.NewVideoRepository(db)
	useCase := usecases.NewVideosUseCase(repo)

	return &VideoHandlers{
		router:  config.GetRouter(),
		useCase: useCase,
	}
}

func (h VideoHandlers) Register() {
	v1 := h.router.Group("/api/v1")
	videos := v1.Group("/videos")
	{
		videos.GET("/", h.index)
		videos.GET("/:id", h.show)
		videos.POST("/", h.create)
		videos.PUT("/:id", h.update)
		videos.DELETE("/:id", h.delete)
	}
}
