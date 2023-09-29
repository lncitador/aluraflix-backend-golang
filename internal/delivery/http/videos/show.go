package videos

import (
	"github.com/gin-gonic/gin"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
)

func (h VideoHandlers) show(c *gin.Context) {
	id := c.Param("id")
	uid, err := v.NewUniqueEntityID(&id)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	video, err := h.useCase.FindById(uid)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, video)
}
