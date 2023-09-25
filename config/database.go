package config

import (
	"github.com/lncitador/alura-flix-backend/internal/domain"
	"gorm.io/gorm"
	"log"
)

func (c Config) GetDb() *gorm.DB {
	return c.db
}

func migrations(c *Config) {
	if err := c.db.AutoMigrate(&domain.Video{}); err != nil {
		log.Fatal(err)
	}
}

func (c Config) Close() {
	sqlDB, err := c.db.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatal(err)
	}

	log.Print("Conexão com o banco de dados encerrada.")
}
