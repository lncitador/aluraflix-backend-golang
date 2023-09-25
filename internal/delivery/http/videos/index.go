package videos

import (
	"github.com/gin-gonic/gin"
)

func (h VideoHandlers) index(c *gin.Context) {
	videos, err := h.useCase.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, videos)
}
