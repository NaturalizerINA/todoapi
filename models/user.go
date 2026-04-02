package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid();column:id" json:"id"`
	Email        string     `gorm:"unique;not null;column:email" json:"email"`
	PasswordHash string     `gorm:"not null;column:password_hash" json:"-"` // Hidden from JSON
	CreatedAt    time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	LastLogin    *time.Time `gorm:"column:last_login" json:"last_login"`
}

func (User) TableName() string {
	return "users"
}

// LoginRequest represents incoming login credentials
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents the structure returned after successful authentication
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
