package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugonorte/chinesonline-go-api/internal/api/handlers"
	"github.com/hugonorte/chinesonline-go-api/internal/api/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	
	// Rotas protegidas pelo Firebase Auth
	sessions := v1.Group("/sessions")
	sessions.Use(middlewares.FirebaseAuthMiddleware())
	{
		sessions.GET("/new", handlers.GenerateSession)
		sessions.POST("/:id/submit", handlers.SubmitSession)
	}

	// Rotas de Usuário (Sync de cadastro)
	users := v1.Group("/users")
	users.Use(middlewares.FirebaseAuthMiddleware())
	{
		users.POST("/sync", handlers.SyncUser)
	}

	// Rotas de Autenticação (Login History)
	auth := v1.Group("/auth")
	auth.Use(middlewares.FirebaseAuthMiddleware())
	{
		auth.POST("/login", handlers.RecordLogin)
	}

	// Rotas administrativas protegidas por Firebase Auth e verificação de Admin
	admin := v1.Group("/admin")
	admin.Use(middlewares.FirebaseAuthMiddleware())
	admin.Use(middlewares.VerifyAdmin())
	{
		admin.POST("/ideograms", handlers.CreateIdeogram)
		admin.PUT("/ideograms/:id", handlers.UpdateIdeogram)
		admin.DELETE("/ideograms/:id", handlers.DeleteIdeogram)
		admin.GET("/ideograms/:id", handlers.GetIdeogram)
		admin.GET("/ideograms", handlers.GetIdeograms)
		admin.POST("/ideograms/batch", handlers.CreateBatchIdeograms)
	}
}
