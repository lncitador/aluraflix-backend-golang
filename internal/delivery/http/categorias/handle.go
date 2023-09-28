package categorias

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories"
	"github.com/lncitador/alura-flix-backend/internal/application/usecases"
)

type CategoriaHandlers struct {
	router  *gin.RouterGroup
	useCase *usecases.CategoriasUseCase
}

func NewCategoriaHandlers(router *gin.RouterGroup, repo repositories.CategoriaRepositoryContract) *CategoriaHandlers {
	useCase := usecases.NewCategoriasUseCase(repo)

	return &CategoriaHandlers{
		router:  router,
		useCase: useCase,
	}
}

func (h CategoriaHandlers) Register(middlewares ...gin.HandlerFunc) {
	categoriasV1 := h.router.Group("/v1/categorias")

	for _, middleware := range middlewares {
		categoriasV1.Use(middleware)
	}

	{
		categoriasV1.GET("/", h.index)
		categoriasV1.GET("/:id", h.show)
		categoriasV1.POST("/", h.create)
		categoriasV1.PUT("/:id", h.update)
		categoriasV1.DELETE("/:id", h.delete)
	}
}
