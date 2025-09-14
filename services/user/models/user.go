package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username     string    `gorm:"size:30;unique;not null"`
	Email        string    `gorm:"size:255;unique;not null"`
	PasswordHash string    `gorm:"not null"`
	Bio          *string
	AvatarURL    *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Optional: before creating, generate UUID if not set
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
