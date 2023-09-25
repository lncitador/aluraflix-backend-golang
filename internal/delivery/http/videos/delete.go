package videos

import (
	"github.com/gin-gonic/gin"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

func (h VideoHandlers) delete(c *gin.Context) {
	id := c.Param("id")
	uid, err := vo.NewUniqueEntityID(&id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.Delete(uid); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}