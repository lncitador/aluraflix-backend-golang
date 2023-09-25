package categorias

import (
	"github.com/gin-gonic/gin"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

func (h CategoriaHandlers) show(c *gin.Context) {
	id := c.Param("id")
	uid, err := v.NewUniqueEntityID(&id)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	categoria, err := h.useCase.FindById(uid)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, categoria)
}
