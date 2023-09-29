package videos

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

func (h VideoHandlers) update(c *gin.Context) {
	id := c.Param("id")
	uid, err := v.NewUniqueEntityID(&id)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	var dto domain.VideoInput
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	video, err := h.useCase.Update(uid, dto)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, video)
}
