package data

import (
	"time"
)

type EmergencyContact struct {
	ID          int64     `gorm:"primaryKey"     json:"id"`
	UserID      int64     `gorm:"not null;index" json:"user_id"`
	FirstName   string    `gorm:"type:text"      json:"first_name"`
	LastName    string    `gorm:"type:text"      json:"last_name"`
	PhoneNumber string    `gorm:"type:text"      json:"phone_number"`
	Email       string    `gorm:"type:text"      json:"email"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type EmergencyContactStore interface {
	CreateEmergencyContact(contact *EmergencyContact) (*EmergencyContact, error)
	GetEmergencyContact(id int64) (*EmergencyContact, error)
	GetEmergencyContactByEmail(email string) (*EmergencyContact, error)
	ListUserEmergencyContacts(userID int64) ([]*EmergencyContact, error)
	UpdateEmergencyContact(contact *EmergencyContact) error
	DeleteEmergencyContact(id int64) error
}
