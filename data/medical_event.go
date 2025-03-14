package data

import (
	"time"
)

type MedicalEvent struct {
	ID          int64     `gorm:"primaryKey"     json:"id"`
	UserID      int64     `gorm:"not null;index" json:"user_id"`
	Age         string    `gorm:"type:text"      json:"age"`
	Description string    `gorm:"type:text"      json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MedicalEventStore interface {
	CreateMedicalEvent(event *MedicalEvent) (*MedicalEvent, error)
	GetMedicalEvent(id int64) (*MedicalEvent, error)
	ListUserMedicalEvents(userID int64) ([]*MedicalEvent, error)
	UpdateMedicalEvent(event *MedicalEvent) error
	DeleteMedicalEvent(id int64) error
}
