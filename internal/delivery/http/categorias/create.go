package categorias

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
)

func (h CategoriaHandlers) create(c *gin.Context) {
	user, _ := c.Get("user")
	var dto domain.CategoriaInput

	dto.UsuarioID = &user.(*domain.UsuarioDto).ID

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	categoria, err := h.useCase.Create(dto)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
	}

	c.JSON(201, categoria)
}
