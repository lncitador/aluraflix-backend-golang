package categorias

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

func (h CategoriaHandlers) update(c *gin.Context) {
	id := c.Param("id")
	uid, err := v.NewUniqueEntityID(&id)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err})
		return
	}

	user, _ := c.Get("user")
	var dto domain.CategoriaInput
	dto.UsuarioID = &user.(*domain.UsuarioDto).ID

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	categoria, err := h.useCase.Update(uid, dto)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err})
		return
	}

	c.JSON(200, categoria)
}
