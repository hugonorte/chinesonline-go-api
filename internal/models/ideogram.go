package models

import (
	"time"
)

type Ideogram struct {
	ID              uint      `gorm:"primaryKey"`
	Character       string    `gorm:"type:varchar(10);not null"`
	Pinyin          string    `gorm:"type:varchar(50)"`
	PinyinWithTones string    `gorm:"type:varchar(50)"`
	Translation     string    `gorm:"type:varchar(255)"`
	DifficultyLevel int       `gorm:"index;default:1"` // 1 a 8
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
