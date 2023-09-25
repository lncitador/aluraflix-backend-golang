package config

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func (c Config) GetRouter() *gin.Engine {
	return c.r
}

func (c Config) RunServe() {
	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta :%s", port)
	if err := c.r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
