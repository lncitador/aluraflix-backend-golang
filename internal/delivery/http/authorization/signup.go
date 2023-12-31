package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
)

func (h AuthHandlers) signup(c *gin.Context) {
	var dto domain.UsuarioInput
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	usuario, err := h.useCase.Create(dto)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, usuario)
}
