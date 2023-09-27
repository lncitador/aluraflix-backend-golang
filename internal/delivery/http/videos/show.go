package videos

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories/errors"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"net/http"
)

func (h VideoHandlers) show(c *gin.Context) {
	id := c.Param("id")
	uid, err := v.NewUniqueEntityID(&id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	video, err := h.useCase.FindById(uid)
	if err != nil {
		if err.Error() == errors.ErrFindByIdVideo {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, video)
}
