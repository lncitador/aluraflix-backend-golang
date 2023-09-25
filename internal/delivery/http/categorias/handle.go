package categorias

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/config"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories/impl"
	"github.com/lncitador/alura-flix-backend/internal/application/usecases"
)

type CategoriaHandlers struct {
	router  *gin.RouterGroup
	useCase *usecases.CategoriasUseCase
}

func NewCategoriaHandlers(config *config.Config) *CategoriaHandlers {
	db := config.GetDb()

	repo := impl.NewCategoriaRepository(db)
	useCase := usecases.NewCategoriasUseCase(repo)

	v1 := config.GetRouter().Group("/api/v1")

	return &CategoriaHandlers{
		router:  v1,
		useCase: useCase,
	}
}

func (h CategoriaHandlers) Register() {
	categorias := h.router.Group("/categorias")
	{
		categorias.GET("/", h.index)
		categorias.GET("/:id", h.show)
	}
}
