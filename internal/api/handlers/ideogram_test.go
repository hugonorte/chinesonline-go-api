package handlers

import (
	"testing"
	"time"

	"github.com/hugonorte/chinesonline-go-api/internal/models"
)

func TestMapToResponse(t *testing.T) {
	now := time.Now()
	idg := models.Ideogram{
		ID:                     42,
		Character:              "你好",
		PinyinWithoutTones:     "ni3 hao3",
		PinyinWithTones:        "nǐ hǎo",
		Translation:            "Olá",
		DifficultyLevel:        2,
		PronunciationAudioFile: "test-audio-123.mp3",
		CreatedAt:              now,
		UpdatedAt:              now,
	}

	cdnBaseURL := "https://cdn.example.com"
	res := mapToResponse(idg, cdnBaseURL)

	if res.ID != idg.ID {
		t.Errorf("expected ID %d, got %d", idg.ID, res.ID)
	}
	if res.Character != idg.Character {
		t.Errorf("expected Character %s, got %s", idg.Character, res.Character)
	}
	if res.Pinyin != idg.PinyinWithoutTones {
		t.Errorf("expected Pinyin %s, got %s", idg.PinyinWithoutTones, res.Pinyin)
	}
	if res.PinyinWithTones != idg.PinyinWithTones {
		t.Errorf("expected PinyinWithTones %s, got %s", idg.PinyinWithTones, res.PinyinWithTones)
	}
	if res.Translation != idg.Translation {
		t.Errorf("expected Translation %s, got %s", idg.Translation, res.Translation)
	}
	if res.DifficultyLevel != idg.DifficultyLevel {
		t.Errorf("expected DifficultyLevel %d, got %d", idg.DifficultyLevel, res.DifficultyLevel)
	}

	expectedAudioURL := "https://cdn.example.com/test-audio-123.mp3"
	if res.AudioURL != expectedAudioURL {
		t.Errorf("expected AudioURL %s, got %s", expectedAudioURL, res.AudioURL)
	}

	if res.CreatedAt != now.Format(time.RFC3339) {
		t.Errorf("expected CreatedAt %s, got %s", now.Format(time.RFC3339), res.CreatedAt)
	}
}

func TestMapToResponseEmptyAudio(t *testing.T) {
	idg := models.Ideogram{
		ID:                     99,
		Character:              "再见",
		PinyinWithoutTones:     "zai4 jian4",
		PinyinWithTones:        "zàijiàn",
		Translation:            "Tchau",
		DifficultyLevel:        3,
		PronunciationAudioFile: "",
	}

	cdnBaseURL := "https://cdn.example.com"
	res := mapToResponse(idg, cdnBaseURL)

	if res.AudioURL != "" {
		t.Errorf("expected empty AudioURL, got %s", res.AudioURL)
	}
}
