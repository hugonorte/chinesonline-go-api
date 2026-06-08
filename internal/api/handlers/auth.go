package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugonorte/chinesonline-go-api/internal/db"
	"github.com/hugonorte/chinesonline-go-api/internal/models"
)

type RecordLoginRequest struct {
	Device string `json:"device"`
}

// RecordLogin insere um registro de histórico de login
func RecordLogin(c *gin.Context) {
	uidData, exists := c.Get("UID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UID não encontrado no token"})
		return
	}
	uid := uidData.(string)

	var req RecordLoginRequest
	// É opcional enviar o JSON, então ignoramos erro se não vier nada
	_ = c.ShouldBindJSON(&req)

	var user models.User
	// Tenta buscar o usuário pelo UID. Se não encontrar, teríamos que recriá-lo
	if err := db.DB.Where(&models.User{UID: uid}).First(&user).Error; err != nil {
		// Usuário não existe na base ainda. Podemos criar um usuário "esqueleto" para salvar o histórico.
		user = models.User{
			UID: uid,
		}
		if err := db.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar usuário implícito: " + err.Error()})
			return
		}
	}

	// Captura os dados do request
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	history := models.LoginHistory{
		UserID:    user.ID,
		IPAddress: ip,
		UserAgent: userAgent,
		Device:    req.Device,
	}

	if err := db.DB.Create(&history).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar histórico de login: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login registrado com sucesso",
	})
}
