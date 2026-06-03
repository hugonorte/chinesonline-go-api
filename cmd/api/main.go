package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hugonorte/chinesonline-go-api/internal/api/middlewares"
	"github.com/hugonorte/chinesonline-go-api/internal/api/routes"
	"github.com/hugonorte/chinesonline-go-api/internal/config"
	"github.com/hugonorte/chinesonline-go-api/internal/db"
)

func main() {
	// Carrega configurações (.env)
	cfg := config.LoadConfig()

	// Inicializa banco de dados
	if cfg.NeonDBUrl != "" {
		db.InitDB(cfg.NeonDBUrl)
	} else {
		log.Println("Aviso: NEON_CONNECTION_STRING não definida. Banco de dados não inicializado.")
	}

	// Inicializa router Gin
	r := gin.Default()

	// Rota de Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Inicializa o Firebase com o arquivo JSON na raiz
	err := middlewares.InitFirebase("chinesonline-prod-firebase-adminsdk-fbsvc-72020d017a.json")
	if err != nil {
		log.Fatalf("Erro ao inicializar o Firebase: %v", err)
	}

	// Registra Rotas
	routes.RegisterRoutes(r)

	log.Printf("Iniciando servidor na porta %s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
