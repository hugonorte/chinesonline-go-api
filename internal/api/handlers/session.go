package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hugonorte/chinesonline-go-api/internal/db"
	"github.com/hugonorte/chinesonline-go-api/internal/models"
)

func generateSalt(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func hashAnswer(answer, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(answer + salt))
	return hex.EncodeToString(hash.Sum(nil))
}

// GenerateSession retorna os ideogramas para o nível solicitado
func GenerateSession(c *gin.Context) {
	level := c.Query("level")
	if level == "" {
		level = "1"
	}

	gameTypeStr := c.Query("game_type")
	if gameTypeStr == "" {
		gameTypeStr = string(models.GameTypePinyinWithoutTone)
	}
	gameType := models.GameType(gameTypeStr)

	uid, exists := c.Get("UID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	// Busca o UserID pelo UID
	var user models.User
	if err := db.DB.Where("uid = ?", uid).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado no banco. Faça login no App."})
		return
	}

	// Busca 10 ideogramas aleatórios até o nível do usuário
	var ideograms []models.Ideogram
	db.DB.Where("difficulty_level <= ?", user.Level).Order("RANDOM()").Limit(10).Find(&ideograms)

	// Prepara a resposta com salt e hash
	var questions []map[string]interface{}
	for _, idg := range ideograms {
		salt := generateSalt(8)
		
		var correctAnswer string
		switch gameType {
		case models.GameTypeTranslation, models.GameTypeTranslationTimed:
			correctAnswer = idg.Translation
		case models.GameTypePinyinWithoutTone, models.GameTypePinyinWithoutToneTimed:
			correctAnswer = idg.PinyinWithoutTones
		case models.GameTypePinyinWithNumericTone, models.GameTypePinyinWithNumericToneTimed:
			correctAnswer = idg.PinyinWithNumericTones
		case models.GameTypePinyinWithSimbolTone, models.GameTypePinyinWithSimbolToneTimed:
			correctAnswer = idg.PinyinWithTones
		default:
			correctAnswer = idg.PinyinWithoutTones
		}

		hash := hashAnswer(correctAnswer, salt)

		questions = append(questions, map[string]interface{}{
			"id":          idg.ID,
			"character":   idg.Character,
			"salt":        salt,
			"hash":        hash,
			"pinyin":      idg.PinyinWithTones,
			"translation": idg.Translation,
		})
	}

	// Cria a sessão
	session := models.QuizSession{
		UserID:         user.ID,
		TimestampStart: time.Now(),
		IsValid:        true,
		GameType:       gameType,
	}
	if err := db.DB.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar sessão"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session_id": session.ID,
		"questions":  questions,
	})
}

type SubmitPayload struct {
	Answers map[string]string `json:"answers"` // Mapeia ID do ideograma para a resposta dada
}

func SubmitSession(c *gin.Context) {
	sessionID := c.Param("id")
	var payload SubmitPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload inválido"})
		return
	}

	var session models.QuizSession
	if err := db.DB.First(&session, "id = ?", sessionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sessão não encontrada"})
		return
	}

	// Validação de Time-Spoofing e Limite Máximo
	timeElapsed := time.Since(session.TimestampStart)
	minTimeRequired := time.Duration(len(payload.Answers)) * time.Second
	
	// Sugestão de limite máximo de segurança: 60 segundos por questão para não ser tão restritivo
	maxTimeAllowed := time.Duration(len(payload.Answers)) * 60 * time.Second
	
	isTimed := session.GameType == models.GameTypeTranslationTimed || 
	           session.GameType == models.GameTypePinyinWithoutToneTimed || 
	           session.GameType == models.GameTypePinyinWithNumericToneTimed || 
	           session.GameType == models.GameTypePinyinWithSimbolToneTimed

	if timeElapsed < minTimeRequired {
		session.IsValid = false
		db.DB.Save(&session)
		c.JSON(http.StatusForbidden, gin.H{"error": "Time spoofing detectado, sessão invalidada."})
		return
	}

	if isTimed && timeElapsed > maxTimeAllowed {
		session.IsValid = false
		db.DB.Save(&session)
		c.JSON(http.StatusForbidden, gin.H{"error": "Tempo máximo excedido, sessão invalidada."})
		return
	}

	score := 0
	for idStr, givenAnswer := range payload.Answers {
		var idg models.Ideogram
		if err := db.DB.First(&idg, idStr).Error; err == nil {
			var expectedAnswer string
			switch session.GameType {
			case models.GameTypeTranslation, models.GameTypeTranslationTimed:
				expectedAnswer = idg.Translation
			case models.GameTypePinyinWithoutTone, models.GameTypePinyinWithoutToneTimed:
				expectedAnswer = idg.PinyinWithoutTones
			case models.GameTypePinyinWithNumericTone, models.GameTypePinyinWithNumericToneTimed:
				expectedAnswer = idg.PinyinWithNumericTones
			case models.GameTypePinyinWithSimbolTone, models.GameTypePinyinWithSimbolToneTimed:
				expectedAnswer = idg.PinyinWithTones
			default:
				expectedAnswer = idg.PinyinWithoutTones
			}

			if expectedAnswer == givenAnswer {
				score += 10
			}
		}
	}

	now := time.Now()
	session.TimestampEnd = &now
	db.DB.Save(&session)

	var user models.User
	db.DB.First(&user, session.UserID)
	
	isNewRecord := false
	if score > user.MaxScore {
		user.MaxScore = score
		isNewRecord = true
		db.DB.Save(&user)
	}

	c.JSON(http.StatusOK, gin.H{
		"score":           score,
		"is_new_record":   isNewRecord,
		"session_valid":   session.IsValid,
	})
}
