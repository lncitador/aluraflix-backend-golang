package categorias

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"net/http"
)

func (h CategoriaHandlers) show(c *gin.Context) {
	id := c.Param("id")
	uid, err := v.NewUniqueEntityID(&id)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userId := user.(*domain.UsuarioDto).ID

	categoria, err := h.useCase.FindById(uid)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	if categoria.UsuarioID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(200, categoria)
}
