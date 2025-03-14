package data

import (
	"time"
)

type FrequentFood struct {
	ID        int64     `gorm:"primaryKey"     json:"id"`
	UserID    int64     `gorm:"not null;index" json:"user_id"`
	FoodName  string    `gorm:"type:text"      json:"food_name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type FrequentFoodStore interface {
	CreateFrequentFood(food *FrequentFood) (*FrequentFood, error)
	GetFrequentFood(id int64) (*FrequentFood, error)
	ListUserFrequentFoods(userID int64) ([]*FrequentFood, error)
	UpdateFrequentFood(food *FrequentFood) error
	DeleteFrequentFood(id int64) error
}
