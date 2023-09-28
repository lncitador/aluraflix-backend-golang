package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	"github.com/lncitador/alura-flix-backend/pkg/validations"
	"net/http"
)

func (h AuthHandlers) signin(c *gin.Context) {
	credentials := struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.useCase.Signin(credentials.Email, credentials.Password)
	if internal, validation := validations.GetErrorsByValidation(err); validation != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": validation})
		return
	} else if internal != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("Authorization", *token, int(domain.TokenMaxAge), "", "", false, true)
	c.JSON(http.StatusCreated, gin.H{})
}
