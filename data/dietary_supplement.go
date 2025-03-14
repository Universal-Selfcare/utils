package data

import (
	"time"
)

type DietarySupplement struct {
	ID        int64     `gorm:"primaryKey"     json:"id"`
	UserID    int64     `gorm:"not null;index" json:"user_id"`
	Name      string    `gorm:"type:text"      json:"name"`
	Dosage    string    `gorm:"type:text"      json:"dosage"`
	StartDate string    `gorm:"type:text"      json:"start_date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type DietarySupplementStore interface {
	CreateDietarySupplement(supplement *DietarySupplement) (*DietarySupplement, error)
	GetDietarySupplement(id int64) (*DietarySupplement, error)
	ListUserDietarySupplements(userID int64) ([]*DietarySupplement, error)
	UpdateDietarySupplement(supplement *DietarySupplement) error
	DeleteDietarySupplement(id int64) error
}
