package data

import (
	"time"
)

type UserIntake struct {
	ID        int64     `gorm:"primaryKey"     json:"id"`
	UserID    int64     `gorm:"not null;index" json:"user_id"`
	FormData  string    `gorm:"type:jsonb"     json:"form_data"` // Store as JSON
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserIntakeStore interface {
	CreateUserIntake(userIntake *UserIntake) (*UserIntake, error)
	GetUserIntakeByUserID(userID int64) (*UserIntake, error)
	UpdateUserIntake(userIntake *UserIntake) error
}
