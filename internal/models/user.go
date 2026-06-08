package models

import (
	"time"
)

type User struct {
	ID          uint        `gorm:"primaryKey"`
	UID         string      `gorm:"uniqueIndex;not null"` // Firebase UID
	Name        string      `gorm:"type:varchar(255)"`
	Email       string      `gorm:"type:varchar(255);uniqueIndex"`
	Status      string      `gorm:"type:varchar(50);default:'active'"`
	Role        string      `gorm:"type:varchar(50);default:'user'"`
	MaxScore    int         `gorm:"default:0"`
	Level       int         `gorm:"default:1"`
	Country     Country     `gorm:"type:smallint;default:0"`
	AccountType AccountType `gorm:"type:smallint;default:0"`
	BirthDate   *time.Time  `gorm:"type:date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
