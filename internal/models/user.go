package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	UID       string    `gorm:"uniqueIndex;not null"` // Firebase UID
	Name      string    `gorm:"type:varchar(255)"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex"`
	Status    string    `gorm:"type:varchar(50);default:'active'"`
	Role      string    `gorm:"type:varchar(50);default:'user'"`
	MaxScore  int       `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
