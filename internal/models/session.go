package models

import "time"

type Session struct {
	ExpiresAt time.Time
	Token     string `gorm:"not null;unique"`
	SessionId string `gorm:"primaryKey;not null"`
}
