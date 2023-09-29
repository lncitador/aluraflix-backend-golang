package videos

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
)

func (h VideoHandlers) index(c *gin.Context) {
	user, _ := c.Get("user")
	userId := user.(*domain.UsuarioDto).ID

	var query domain.VideoQuery
	if err := query.SetUsuarioID(userId); err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	search := c.Query("search")
	query.SetSearch(search)

	page := c.Query("page")
	if err := query.SetPage(page); err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	videos, err := h.useCase.FindAll(&query)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	if query.Page() != nil {
		pagination := domain.Pagination[domain.VideoDto]{}

		if err := pagination.Paginate(c.Request.URL.String(), videos, &query); err != nil {
			c.JSON(err.Status(), gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, pagination)
		return
	}

	c.JSON(200, videos)
}
