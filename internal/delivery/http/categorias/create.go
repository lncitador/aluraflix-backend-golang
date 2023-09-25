package categorias

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
)

func (h CategoriaHandlers) create(c *gin.Context) {
	var dto domain.CategoriaInput
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	categoria, err := h.useCase.Create(dto)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, categoria)
}
