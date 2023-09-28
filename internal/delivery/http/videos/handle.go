package videos

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories"
	"github.com/lncitador/alura-flix-backend/internal/application/usecases"
)

type VideoHandlers struct {
	router  *gin.RouterGroup
	useCase *usecases.VideosUseCase
}

func NewVideoHandlers(router *gin.RouterGroup, repo repositories.VideoRepositoryContract) *VideoHandlers {
	useCase := usecases.NewVideosUseCase(repo)

	return &VideoHandlers{
		router:  router,
		useCase: useCase,
	}
}

func (h VideoHandlers) Register() {
	videosV1 := h.router.Group("/v1/videos")
	{
		videosV1.GET("/", h.index)
		videosV1.GET("/:id", h.show)
		videosV1.POST("/", h.create)
		videosV1.PUT("/:id", h.update)
		videosV1.DELETE("/:id", h.delete)
	}
}
