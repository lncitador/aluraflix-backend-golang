package categorias

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	"net/http"
)

func (h CategoriaHandlers) index(c *gin.Context) {
	user, _ := c.Get("user")
	var query domain.CategoriaQuery

	if err := query.SetUsuarioID(user.(*domain.UsuarioDto).ID); err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	categorias, err := h.useCase.FindAll(&query)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categorias)
}
