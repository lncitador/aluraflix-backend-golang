package categorias

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	"github.com/lncitador/alura-flix-backend/pkg/validations"
	"net/http"
)

func (h CategoriaHandlers) create(c *gin.Context) {
	var dto domain.CategoriaInput
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	categoria, err := h.useCase.Create(dto)
	if internal, validation := validations.GetErrorsByValidation(err); validation != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": validation})
		return
	} else if internal != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, categoria)
}
