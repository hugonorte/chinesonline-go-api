package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hugonorte/chinesonline-go-api/internal/db"
	"github.com/hugonorte/chinesonline-go-api/internal/models"
)

type SyncUserRequest struct {
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Country     int        `json:"country"`
	AccountType int        `json:"account_type"`
	BirthDate   *time.Time `json:"birth_date" time_format:"2006-01-02"`
}

// SyncUser realiza um UPSERT do usuário usando o UID extraído do JWT
func SyncUser(c *gin.Context) {
	uidData, exists := c.Get("UID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UID não encontrado no token"})
		return
	}
	uid := uidData.(string)

	var req SyncUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de dados inválido: " + err.Error()})
		return
	}

	var user models.User
	// Tenta buscar o usuário pelo UID
	result := db.DB.Where(&models.User{UID: uid}).First(&user)

	// Prepara os dados baseados no payload
	user.UID = uid
	user.Name = req.Name
	user.Email = req.Email
	user.Country = models.Country(req.Country)
	user.AccountType = models.AccountType(req.AccountType)
	user.BirthDate = req.BirthDate

	if result.Error != nil {
		// Usuário não existe, vamos criar
		if err := db.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário: " + err.Error()})
			return
		}
	} else {
		// Usuário já existe, vamos atualizar
		if err := db.DB.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário sincronizado com sucesso",
		"user":    user,
	})
}

// GetMe retorna o perfil do usuário logado
func GetMe(c *gin.Context) {
	uidData, exists := c.Get("UID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UID não encontrado no token"})
		return
	}
	uid := uidData.(string)

	var user models.User
	if err := db.DB.Where(&models.User{UID: uid}).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado. Faça o sync primeiro."})
		return
	}

	c.JSON(http.StatusOK, user)
}
