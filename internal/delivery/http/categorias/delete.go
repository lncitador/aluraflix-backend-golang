package categorias

import (
	"github.com/gin-gonic/gin"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"github.com/lncitador/alura-flix-backend/pkg/validations"
	"net/http"
)

func (h CategoriaHandlers) delete(c *gin.Context) {
	id := c.Param("id")
	uid, err := v.NewUniqueEntityID(&id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.useCase.Delete(uid)
	if internal, validation := validations.GetErrorsByValidation(err); validation != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": validation})
		return
	} else if internal != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": internal})
		return
	}

	c.JSON(204, nil)
}
