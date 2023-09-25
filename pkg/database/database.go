package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

// Inicializar inicializa a conexão com o banco de dados.
func Inicializar() (*gorm.DB, error) {
	var err error

	// Configurar o driver de banco de dados e a conexão.
	// Substitua pelo driver de banco de dados que você está usando (ex: Postgres, MySQL, etc.).
	url := os.Getenv("DATABASE_URL")
	db, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
		return nil, err
	}

	return db, nil
}
