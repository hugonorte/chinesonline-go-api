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
}
