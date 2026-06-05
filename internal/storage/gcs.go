package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

// UploadAudio uploads an audio file to the configured GCS bucket.
// It generates a unique name using UUID, uploads it, and returns the generated filename.
func UploadAudio(ctx context.Context, reader io.Reader, originalFilename string) (string, error) {
	bucketName := os.Getenv("GCS_BUCKET_NAME")
	if bucketName == "" {
		return "", fmt.Errorf("GCS_BUCKET_NAME não configurado no ambiente")
	}

	// Criar o cliente GCS
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("falha ao criar cliente GCS: %w", err)
	}
	defer client.Close()

	// Determinar a extensão e o Content-Type apropriado
	ext := strings.ToLower(filepath.Ext(originalFilename))
	if ext == "" {
		ext = ".mp3" // Default para mp3
	}

	var contentType string
	switch ext {
	case ".mp3":
		contentType = "audio/mpeg"
	case ".wav":
		contentType = "audio/wav"
	case ".ogg":
		contentType = "audio/ogg"
	case ".m4a":
		contentType = "audio/mp4"
	default:
		contentType = "application/octet-stream"
	}

	// Gerar um nome único para o arquivo
	uniqueFilename := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// Fazer upload para o bucket
	bucket := client.Bucket(bucketName)
	object := bucket.Object(uniqueFilename)

	wc := object.NewWriter(ctx)
	wc.ContentType = contentType
	// Como o bucket já está configurado como público (allUsers -> Storage Object Viewer) via IAM uniforme,
	// não precisamos passar ACL explicitamente por objeto. Mas garantimos os metadados corretos.

	if _, err := io.Copy(wc, reader); err != nil {
		wc.Close()
		return "", fmt.Errorf("falha ao copiar arquivo para o GCS writer: %w", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("falha ao fechar GCS writer: %w", err)
	}

	return uniqueFilename, nil
}

// DeleteAudio deletes an audio file from the GCS bucket.
func DeleteAudio(ctx context.Context, filename string) error {
	if filename == "" {
		return nil
	}

	bucketName := os.Getenv("GCS_BUCKET_NAME")
	if bucketName == "" {
		return fmt.Errorf("GCS_BUCKET_NAME não configurado no ambiente")
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("falha ao criar cliente GCS: %w", err)
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	object := bucket.Object(filename)

	if err := object.Delete(ctx); err != nil {
		// Se o arquivo não existir, não tratamos como erro crítico ao deletar
		if err == storage.ErrObjectNotExist {
			return nil
		}
		return fmt.Errorf("falha ao deletar arquivo do GCS: %w", err)
	}

	return nil
}
