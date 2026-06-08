package models

import (
	"time"
)

type QuizSession struct {
	ID             string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // PostgreSQL uuid
	UserID         uint      `gorm:"index;not null"`
	TimestampStart time.Time `gorm:"not null"`
	TimestampEnd   *time.Time
	IsValid        bool      `gorm:"default:true"`
	GameType       GameType  `gorm:"type:varchar(50);default:'pinyin_without_tone'"`
}
