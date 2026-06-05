package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hugonorte/chinesonline-go-api/internal/db"
	"github.com/hugonorte/chinesonline-go-api/internal/models"
	"github.com/hugonorte/chinesonline-go-api/internal/storage"
)

// IdeogramResponse define a estrutura de resposta JSON compatível com o Frontend Vue/Nuxt
type IdeogramResponse struct {
	ID              uint   `json:"id"`
	Character       string `json:"character"`
	Pinyin          string `json:"pinyin"`
	PinyinWithTones string `json:"pinyin_with_tones"`
	Translation     string `json:"translation"`
	DifficultyLevel int    `json:"difficulty_level"`
	AudioURL        string `json:"audio_url,omitempty"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

// BatchIdeogramInput define o formato para cadastros em lote recebidos em JSON
type BatchIdeogramInput struct {
	Character       string `json:"character" binding:"required"`
	Pinyin          string `json:"pinyin" binding:"required"`
	PinyinWithTones string `json:"pinyin_with_tones" binding:"required"`
	Translation     string `json:"translation" binding:"required"`
	DifficultyLevel int    `json:"difficulty_level" binding:"required"`
}

// Auxiliar para mapear o modelo de banco para a estrutura de resposta da API
func mapToResponse(idg models.Ideogram, cdnBaseURL string) IdeogramResponse {
	audioURL := ""
	if idg.PronunciationAudioFile != "" && cdnBaseURL != "" {
		audioURL = fmt.Sprintf("%s/%s", cdnBaseURL, idg.PronunciationAudioFile)
	}
	return IdeogramResponse{
		ID:              idg.ID,
		Character:       idg.Character,
		Pinyin:          idg.PinyinWithoutTones,
		PinyinWithTones: idg.PinyinWithTones,
		Translation:     idg.Translation,
		DifficultyLevel: idg.DifficultyLevel,
		AudioURL:        audioURL,
		CreatedAt:       idg.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       idg.UpdatedAt.Format(time.RFC3339),
	}
}

// CreateIdeogram cria um novo ideograma a partir de um formulário multipart/form-data com upload de áudio
func CreateIdeogram(c *gin.Context) {
	character := c.PostForm("character")
	pinyin := c.PostForm("pinyin")
	pinyinWithTones := c.PostForm("pinyin_with_tones")
	translation := c.PostForm("translation")
	difficultyLevelStr := c.PostForm("difficulty_level")

	if character == "" || pinyin == "" || pinyinWithTones == "" || translation == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos os campos de texto obrigatórios (character, pinyin, pinyin_with_tones, translation) devem ser preenchidos"})
		return
	}

	difficultyLevel := 1
	if difficultyLevelStr != "" {
		if dl, err := strconv.Atoi(difficultyLevelStr); err == nil {
			difficultyLevel = dl
		}
	}

	var audioFilename string
	audioFile, header, err := c.Request.FormFile("audio")
	if err == nil {
		defer audioFile.Close()
		
		ext := strings.ToLower(filepath.Ext(header.Filename))
		if ext != ".mp3" && ext != ".wav" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de arquivo de áudio inválido. Apenas .mp3 e .wav são permitidos."})
			return
		}

		if header.Size > 500*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "O arquivo de áudio deve ter no máximo 500KB."})
			return
		}

		var uploadErr error
		audioFilename, uploadErr = storage.UploadAudio(c.Request.Context(), audioFile, header.Filename)
		if uploadErr != nil {
			log.Printf("Erro ao fazer upload de áudio no GCS: %v", uploadErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao salvar arquivo de áudio: %v", uploadErr)})
			return
		}
	}

	ideogram := models.Ideogram{
		Character:              character,
		PinyinWithoutTones:     pinyin,
		PinyinWithTones:        pinyinWithTones,
		Translation:            translation,
		DifficultyLevel:        difficultyLevel,
		PronunciationAudioFile: audioFilename,
	}

	if err := db.DB.Create(&ideogram).Error; err != nil {
		// Se falhou ao salvar no banco, tentamos remover o arquivo que foi carregado no Storage
		if audioFilename != "" {
			_ = storage.DeleteAudio(c.Request.Context(), audioFilename)
		}
		log.Printf("Erro ao criar ideograma no banco: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar ideograma no banco de dados"})
		return
	}

	cdnBaseURL := os.Getenv("CDN_BASE_URL")
	c.JSON(http.StatusCreated, mapToResponse(ideogram, cdnBaseURL))
}

// UpdateIdeogram atualiza parcialmente ou totalmente um ideograma (via multipart/form-data)
func UpdateIdeogram(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do ideograma inválido"})
		return
	}

	var ideogram models.Ideogram
	if err := db.DB.First(&ideogram, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ideograma não encontrado"})
		return
	}

	character := c.PostForm("character")
	pinyin := c.PostForm("pinyin")
	pinyinWithTones := c.PostForm("pinyin_with_tones")
	translation := c.PostForm("translation")
	difficultyLevelStr := c.PostForm("difficulty_level")

	if character != "" {
		ideogram.Character = character
	}
	if pinyin != "" {
		ideogram.PinyinWithoutTones = pinyin
	}
	if pinyinWithTones != "" {
		ideogram.PinyinWithTones = pinyinWithTones
	}
	if translation != "" {
		ideogram.Translation = translation
	}
	if difficultyLevelStr != "" {
		if dl, err := strconv.Atoi(difficultyLevelStr); err == nil {
			ideogram.DifficultyLevel = dl
		}
	}

	// Tratamento do áudio caso tenha sido enviado um novo arquivo
	audioFile, header, err := c.Request.FormFile("audio")
	if err == nil {
		defer audioFile.Close()

		ext := strings.ToLower(filepath.Ext(header.Filename))
		if ext != ".mp3" && ext != ".wav" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de arquivo de áudio inválido. Apenas .mp3 e .wav são permitidos."})
			return
		}

		if header.Size > 500*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "O arquivo de áudio deve ter no máximo 500KB."})
			return
		}

		// Guarda o áudio antigo para deletar depois do sucesso
		oldAudioFilename := ideogram.PronunciationAudioFile

		// Faz upload do novo áudio
		newAudioFilename, uploadErr := storage.UploadAudio(c.Request.Context(), audioFile, header.Filename)
		if uploadErr != nil {
			log.Printf("Erro ao fazer upload do novo áudio no GCS: %v", uploadErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao salvar novo arquivo de áudio: %v", uploadErr)})
			return
		}

		ideogram.PronunciationAudioFile = newAudioFilename

		// Atualiza no banco
		if err := db.DB.Save(&ideogram).Error; err != nil {
			// Se falhar no banco, deleta o novo arquivo recém-criado do storage
			_ = storage.DeleteAudio(c.Request.Context(), newAudioFilename)
			log.Printf("Erro ao atualizar ideograma no banco: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar ideograma no banco de dados"})
			return
		}

		// Se atualizou com sucesso e existia um áudio anterior, exclui o anterior do GCS
		if oldAudioFilename != "" {
			if delErr := storage.DeleteAudio(c.Request.Context(), oldAudioFilename); delErr != nil {
				log.Printf("Aviso: falha ao deletar áudio antigo %s do GCS: %v", oldAudioFilename, delErr)
			}
		}
	} else {
		// Nenhuma alteração de áudio, apenas salva os campos de texto no banco
		if err := db.DB.Save(&ideogram).Error; err != nil {
			log.Printf("Erro ao atualizar ideograma no banco: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar ideograma no banco de dados"})
			return
		}
	}

	cdnBaseURL := os.Getenv("CDN_BASE_URL")
	c.JSON(http.StatusOK, mapToResponse(ideogram, cdnBaseURL))
}

// DeleteIdeogram remove um ideograma do banco de dados e limpa seu respectivo arquivo no storage GCS
func DeleteIdeogram(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do ideograma inválido"})
		return
	}

	var ideogram models.Ideogram
	if err := db.DB.First(&ideogram, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ideograma não encontrado"})
		return
	}

	// Remove o áudio associado se houver
	if ideogram.PronunciationAudioFile != "" {
		if delErr := storage.DeleteAudio(c.Request.Context(), ideogram.PronunciationAudioFile); delErr != nil {
			log.Printf("Aviso: falha ao deletar áudio %s do GCS na exclusão: %v", ideogram.PronunciationAudioFile, delErr)
		}
	}

	if err := db.DB.Delete(&ideogram).Error; err != nil {
		log.Printf("Erro ao deletar ideograma no banco: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir ideograma do banco de dados"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ideograma excluído com sucesso"})
}

// GetIdeogram retorna os dados de um único ideograma mapeando o arquivo de áudio para a URL do CDN
func GetIdeogram(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do ideograma inválido"})
		return
	}

	var ideogram models.Ideogram
	if err := db.DB.First(&ideogram, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ideograma não encontrado"})
		return
	}

	cdnBaseURL := os.Getenv("CDN_BASE_URL")
	c.JSON(http.StatusOK, mapToResponse(ideogram, cdnBaseURL))
}

// GetIdeograms retorna a lista paginada e filtrada de todos os ideogramas cadastrados
func GetIdeograms(c *gin.Context) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	levelStr := c.Query("difficulty_level")

	page := 1
	limit := 100 // Padrão alinhado ao comportamento do frontend

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	offset := (page - 1) * limit

	query := db.DB.Model(&models.Ideogram{})
	if levelStr != "" {
		if level, err := strconv.Atoi(levelStr); err == nil {
			query = query.Where("difficulty_level = ?", level)
		}
	}

	var ideograms []models.Ideogram
	if err := query.Order("id asc").Offset(offset).Limit(limit).Find(&ideograms).Error; err != nil {
		log.Printf("Erro ao listar ideogramas: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar ideogramas do banco"})
		return
	}

	cdnBaseURL := os.Getenv("CDN_BASE_URL")
	responses := make([]IdeogramResponse, len(ideograms))
	for i, idg := range ideograms {
		responses[i] = mapToResponse(idg, cdnBaseURL)
	}

	// Retorna o array diretamente conforme interface do Nuxt (BackendIdeograma[])
	c.JSON(http.StatusOK, responses)
}

// CreateBatchIdeograms cadastra múltiplos ideogramas a partir de uma lista JSON (sem envio de áudio)
func CreateBatchIdeograms(c *gin.Context) {
	var inputs []BatchIdeogramInput
	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload JSON inválido: " + err.Error()})
		return
	}

	if len(inputs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nenhum ideograma fornecido no lote"})
		return
	}

	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var ideograms []models.Ideogram
	for idx, input := range inputs {
		if input.Character == "" || input.Pinyin == "" || input.PinyinWithTones == "" || input.Translation == "" {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Campos obrigatórios ausentes no item de índice %d (character, pinyin, pinyin_with_tones e translation)", idx)})
			return
		}

		ideograms = append(ideograms, models.Ideogram{
			Character:          input.Character,
			PinyinWithoutTones: input.Pinyin,
			PinyinWithTones:    input.PinyinWithTones,
			Translation:        input.Translation,
			DifficultyLevel:    input.DifficultyLevel,
		})
	}

	if err := tx.Create(&ideograms).Error; err != nil {
		tx.Rollback()
		log.Printf("Erro ao inserir lote no banco: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar ideogramas em lote no banco de dados"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Erro ao fazer commit da transação de lote: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao finalizar cadastro em lote"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d ideogramas cadastrados com sucesso", len(ideograms)),
		"count":   len(ideograms),
	})
}
