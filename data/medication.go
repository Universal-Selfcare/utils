package data

import (
	"time"
)

type Medication struct {
	ID          int64     `gorm:"primaryKey"     json:"id"`
	UserID      int64     `gorm:"not null;index" json:"user_id"`
	Name        string    `gorm:"type:text"      json:"name"`
	Dosage      string    `gorm:"type:text"      json:"dosage"`
	StartDate   time.Time `                      json:"start_date"`
	EndDate     time.Time `                      json:"end_date"`
	Current     bool      `                      json:"current"`
	SideEffects string    `gorm:"type:text"      json:"side_effects"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MedicationStore interface {
	CreateMedication(medication *Medication) (*Medication, error)
	GetMedication(id int64) (*Medication, error)
	ListUserMedications(userID int64) ([]*Medication, error)
	ListUserCurrentMedications(userID int64) ([]*Medication, error)
	UpdateMedication(medication *Medication) error
	DeleteMedication(id int64) error
}
