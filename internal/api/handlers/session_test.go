package handlers

import (
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	salt1 := generateSalt(8)
	salt2 := generateSalt(8)

	if len(salt1) != 8 || len(salt2) != 8 {
		t.Errorf("Expected salt length 8, got %d and %d", len(salt1), len(salt2))
	}

	if salt1 == salt2 {
		t.Errorf("Expected different salts, got same %s", salt1)
	}
}

func TestHashAnswer(t *testing.T) {
	answer := "Nihao"
	salt := "abcdef"

	hash1 := hashAnswer(answer, salt)
	hash2 := hashAnswer(answer, salt)

	if hash1 != hash2 {
		t.Errorf("Expected same hash for same inputs, got %s and %s", hash1, hash2)
	}

	if hash1 == "" {
		t.Errorf("Expected non-empty hash")
	}
}
