package data

import (
	"time"
)

type Allergy struct {
	ID          int64     `gorm:"primaryKey"     json:"id"`
	UserID      int64     `gorm:"not null;index" json:"user_id"`
	AllergyName string    `gorm:"type:text"      json:"allergy_name"`
	Reaction    string    `gorm:"type:text"      json:"reaction"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type AllergyStore interface {
	CreateAllergy(allergy *Allergy) (*Allergy, error)
	GetAllergy(id int64) (*Allergy, error)
	ListUserAllergies(userID int64) ([]*Allergy, error)
	UpdateAllergy(allergy *Allergy) error
	DeleteAllergy(id int64) error
}
