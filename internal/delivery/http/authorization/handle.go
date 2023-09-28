package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories"
	"github.com/lncitador/alura-flix-backend/internal/application/usecases"
)

type AuthHandlers struct {
	router  *gin.RouterGroup
	useCase *usecases.UsuariosUseCase
}

func NewAuthHandlers(router *gin.RouterGroup, repo repositories.UsuarioRepositoryContract) *AuthHandlers {
	useCase := usecases.NewUsuariosUseCase(repo)

	return &AuthHandlers{
		router:  router,
		useCase: useCase,
	}
}

func (h AuthHandlers) Register() {
	authV1 := h.router.Group("/v1/auth")
	{
		authV1.POST("/signin", h.signin)
		authV1.POST("/signup", h.signup)
	}
}
