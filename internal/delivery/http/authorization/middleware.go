package authorization

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h AuthHandlers) AuthMiddleware(c *gin.Context) {
	fmt.Println("AuthMiddleware")
	c.Next()
}
