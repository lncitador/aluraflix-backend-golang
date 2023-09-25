package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lncitador/alura-flix-backend/pkg/database"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	db *gorm.DB
	r  *gin.Engine
}

func NewConfig() *Config {
	config := &Config{
		r: gin.Default(),
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	fmt.Print("Configurando banco de dados...\n")
	db, err := database.Inicializar()
	if err != nil {
		log.Fatal(err)
	}

	config.db = db

	fmt.Print("Configurando migrações...\n")
	migrations(config)

	return config
}
