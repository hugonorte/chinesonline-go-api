package models

import (
	"time"
)

type LoginHistory struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"`       // Referência ao usuário
	IPAddress string    `gorm:"type:varchar(45)"`     // Suporta IPv4 e IPv6
	UserAgent string    `gorm:"type:text"`            // Para saber o navegador/OS (ex: Chrome no Windows)
	Device    string    `gorm:"type:varchar(100)"`    // (Opcional) ex: "Mobile", "Desktop", "iOS", "Android"
	LoginTime time.Time `gorm:"autoCreateTime"`       // Data e hora exata do login
	IsSuccess bool      `gorm:"default:true"`         // Caso queira registrar tentativas falhas de login no futuro
}
