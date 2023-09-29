package videos

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
)

func (h VideoHandlers) create(c *gin.Context) {
	user, _ := c.Get("userId")
	var dto domain.VideoInput
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	dto.UsuarioID = &user.(*domain.UsuarioDto).ID

	video, err := h.useCase.Create(dto)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err})
		return
	}

	c.JSON(201, video)
}
