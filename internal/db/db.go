package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/hugonorte/chinesonline-go-api/internal/models"
)

var DB *gorm.DB

func InitDB(dsn string) {
	var err error

	// Connect to Neon serverless Postgres
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Falha ao conectar no banco de dados Neon: %v", err)
	}

	// Auto-Migrate
	log.Println("Rodando AutoMigrate...")
	DB.AutoMigrate(&models.User{}, &models.Ideogram{}, &models.QuizSession{})

	log.Println("Conectado ao Neon Database com sucesso!")
}
