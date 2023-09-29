package videos

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

func (h VideoHandlers) show(c *gin.Context) {
	id := c.Param("id")
	uid, err := v.NewUniqueEntityID(&id)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err})
		return
	}

	user, _ := c.Get("user")
	id = user.(*domain.UsuarioDto).ID
	userId, err := v.NewUniqueEntityID(&id)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err})
		return
	}

	video, err := h.useCase.FindById(uid, userId)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err})
		return
	}

	c.JSON(200, video)
}
