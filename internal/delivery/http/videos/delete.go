package videos

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

func (h VideoHandlers) delete(c *gin.Context) {
	id := c.Param("id")
	uid, err := vo.NewUniqueEntityID(&id)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err})
		return
	}

	user, _ := c.Get("user")
	id = user.(*domain.UsuarioDto).ID
	userId, err := vo.NewUniqueEntityID(&id)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err})
		return
	}

	if err := h.useCase.Delete(uid, userId); err != nil {
		c.JSON(err.Status(), gin.H{"error": err})
		return
	}

	c.JSON(204, nil)
}
