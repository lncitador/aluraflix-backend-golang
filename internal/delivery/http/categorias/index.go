package categorias

import "github.com/gin-gonic/gin"

func (h CategoriaHandlers) index(c *gin.Context) {
	categorias, err := h.useCase.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, categorias)
}
