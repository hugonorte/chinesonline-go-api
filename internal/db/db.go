package db

import (
	"log"
	"os"

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

	// Seed para Testes de Desenvolvimento
	if os.Getenv("APP_ENV") == "development" {
		var user models.User
		DB.FirstOrCreate(&user, models.User{UID: "admin-local-123", Name: "Admin Test", Email: "admin@test.com", Role: "admin"})

		var count int64
		DB.Model(&models.Ideogram{}).Count(&count)
		if count == 0 {
			DB.Create([]models.Ideogram{
				{Character: "我", Pinyin: "wo3", PinyinWithTones: "wǒ", Translation: "Eu", DifficultyLevel: 1},
				{Character: "你", Pinyin: "ni3", PinyinWithTones: "nǐ", Translation: "Você", DifficultyLevel: 1},
				{Character: "好", Pinyin: "hao3", PinyinWithTones: "hǎo", Translation: "Bom / Bem", DifficultyLevel: 1},
			})
			log.Println("Ideogramas de teste (seed) inseridos com sucesso!")
		}
	}

	log.Println("Conectado ao Neon Database com sucesso!")
}
