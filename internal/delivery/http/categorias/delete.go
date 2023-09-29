package categorias

import (
	"github.com/gin-gonic/gin"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

func (h CategoriaHandlers) delete(c *gin.Context) {
	id := c.Param("id")
	uid, err := v.NewUniqueEntityID(&id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.Delete(uid); err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}
