package videos

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
)

func (h VideoHandlers) index(c *gin.Context) {
	var query domain.VideoQuery
	search := c.Query("search")
	query.SetSearch(search)

	page := c.Query("page")
	if err := query.SetPage(page); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	videos, err := h.useCase.FindAll(query)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, videos)
}
