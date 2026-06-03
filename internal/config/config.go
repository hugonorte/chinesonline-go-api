package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	NeonDBUrl string
	Port      string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: arquivo .env não encontrado. Utilizando variáveis de ambiente do sistema.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Use NEON_CONNECTION_STRING from .env
	neonURL := os.Getenv("NEON_CONNECTION_STRING")

	return &Config{
		NeonDBUrl: neonURL,
		Port:      port,
	}
}
