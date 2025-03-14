package data

import (
	"time"
)

type Caregiver struct {
	ID          int64     `gorm:"primaryKey"                json:"id"`
	UserID      int64     `gorm:"not null;index"            json:"user_id"`
	Email       string    `gorm:"type:text;not null;index"  json:"email"`
	PhoneNumber string    `gorm:"type:text;not null;index;" json:"phone_number"`
	CreatedAt   time.Time `gorm:"autoCreateTime"            json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"            json:"updated_at"`
}

type CaregiverStore interface {
	CreateCaregiver(caregiver *Caregiver) (*Caregiver, error)
	GetCaregiver(id int64) (*Caregiver, error)
	GetCaregiverByEmail(email string) (*Caregiver, error)
	ListUserCaregivers(userID int64) ([]*Caregiver, error)
	UpdateCaregiver(caregiver *Caregiver) error
	DeleteCaregiver(id int64) error
}
