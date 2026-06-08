package models

import (
	"time"
)

type Ideogram struct {
	ID                     uint      `gorm:"primaryKey"`
	Character              string    `gorm:"type:varchar(10);not null"`
	PinyinWithTones        string    `gorm:"type:varchar(50)"`
	PinyinWithoutTones     string    `gorm:"type:varchar(50)"`
	PinyinWithNumericTones string    `gorm:"type:varchar(50)"`
	Translation            string    `gorm:"type:varchar(255)"`
	DifficultyLevel        int       `gorm:"index;default:1"` // 1 a 8
	PronunciationAudioFile string    `gorm:"type:varchar(255)"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
