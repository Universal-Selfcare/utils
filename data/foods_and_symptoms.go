package data

import (
	"time"
)

// TrackingPeriod represents a 5-day tracking period for a user
type TrackingPeriod struct {
	ID          int64     `gorm:"primaryKey"     json:"id"`
	UserID      int64     `gorm:"not null;index" json:"user_id"`
	StartDate   time.Time `gorm:"not null"       json:"start_date"`
	EndDate     time.Time `gorm:"not null"       json:"end_date"`
	IsCompleted bool      `gorm:"default:false"  json:"is_completed"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// MealEntry represents a single meal entry during a tracking day
type MealEntry struct {
	ID               int64     `gorm:"primaryKey"     json:"id"`
	UserID           int64     `gorm:"not null;index" json:"user_id"`
	TrackingPeriodID int64     `gorm:"not null;index" json:"tracking_period_id"`
	TrackingDay      int       `gorm:"not null"       json:"tracking_day"` // 1-5
	MealType         string    `gorm:"not null"       json:"meal_type"`    // Breakfast, Lunch, Snack, Dinner
	MealTime         string    `gorm:"not null"       json:"meal_time"`    // Time the meal was eaten
	MealDuration     string    `gorm:"not null"       json:"meal_duration"`
	Notes            string    `gorm:"type:text"      json:"notes"`
	PortionSize      string    `gorm:"not null"       json:"portion_size"`
	IsCompleted      bool      `gorm:"default:false"  json:"is_completed"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// FoodItem represents a food item from the predefined list
type FoodItem struct {
	ID        int64     `gorm:"primaryKey"     json:"id"`
	Name      string    `gorm:"not null;index" json:"name"`
	Category  string    `gorm:"not null"       json:"category"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// MealFood links a meal entry to a food item
type MealFood struct {
	ID          int64     `gorm:"primaryKey"     json:"id"`
	MealEntryID int64     `gorm:"not null;index" json:"meal_entry_id"`
	FoodItemID  int64     `gorm:"not null"       json:"food_item_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// CustomFood represents a custom food added by a user
type CustomFood struct {
	ID          int64     `gorm:"primaryKey"     json:"id"`
	MealEntryID int64     `gorm:"not null;index" json:"meal_entry_id"`
	Name        string    `gorm:"not null"       json:"name"`
	Portion     string    `gorm:"not null"       json:"portion"`
	Preparation string    `gorm:"not null"       json:"preparation"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// Symptom represents a tracked symptom related to a meal
type Symptom struct {
	ID          int64     `gorm:"primaryKey"     json:"id"`
	MealEntryID int64     `gorm:"not null;index" json:"meal_entry_id"`
	SymptomType string    `gorm:"not null"       json:"symptom_type"`
	Severity    int       `gorm:"not null"       json:"severity"` // 0-180
	IsOvernight bool      `gorm:"default:false"  json:"is_overnight"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TrackingPeriodStore provides database operations for tracking periods
type TrackingPeriodStore interface {
	CreateTrackingPeriod(period *TrackingPeriod) (*TrackingPeriod, error)
	GetTrackingPeriod(id int64) (*TrackingPeriod, error)
	GetCurrentTrackingPeriod(userID int64) (*TrackingPeriod, error)
	GetLastCompletedTrackingPeriod(userID int64) (*TrackingPeriod, error)
	ListUserTrackingPeriods(userID int64) ([]*TrackingPeriod, error)
	UpdateTrackingPeriod(period *TrackingPeriod) error
	CompleteTrackingPeriod(id int64) error
}

// MealEntryStore provides database operations for meal entries
type MealEntryStore interface {
	CreateMealEntry(entry *MealEntry) (*MealEntry, error)
	GetMealEntry(id int64) (*MealEntry, error)
	GetMealEntryByDetails(userID int64, trackingPeriodID int64, day int, mealType string) (*MealEntry, error)
	ListUserMealEntries(userID int64, trackingPeriodID int64) ([]*MealEntry, error)
	UpdateMealEntry(entry *MealEntry) error
	CompleteMealEntry(id int64) error
	DeleteMealEntry(id int64) error
}

// FoodItemStore provides database operations for food items
type FoodItemStore interface {
	CreateFoodItem(item *FoodItem) (*FoodItem, error)
	GetFoodItem(id int64) (*FoodItem, error)
	GetFoodItemByName(name string) (*FoodItem, error)
	ListFoodItems() ([]*FoodItem, error)
	ListFoodItemsByCategory(category string) ([]*FoodItem, error)
	UpdateFoodItem(item *FoodItem) error
	DeleteFoodItem(id int64) error
}

// MealFoodStore provides database operations for meal foods
type MealFoodStore interface {
	CreateMealFood(mealFood *MealFood) (*MealFood, error)
	GetMealFoodsForMeal(mealEntryID int64) ([]*MealFood, error)
	DeleteMealFood(id int64) error
	DeleteAllMealFoodsForMeal(mealEntryID int64) error
}

// CustomFoodStore provides database operations for custom foods
type CustomFoodStore interface {
	CreateCustomFood(food *CustomFood) (*CustomFood, error)
	GetCustomFoodsForMeal(mealEntryID int64) ([]*CustomFood, error)
	UpdateCustomFood(food *CustomFood) error
	DeleteCustomFood(id int64) error
	DeleteAllCustomFoodsForMeal(mealEntryID int64) error
}

// SymptomStore provides database operations for symptoms
type SymptomStore interface {
	CreateSymptom(symptom *Symptom) (*Symptom, error)
	GetSymptom(id int64) (*Symptom, error)
	GetSymptomByTypeForMeal(mealEntryID int64, symptomType string) (*Symptom, error)
	ListSymptomsForMeal(mealEntryID int64) ([]*Symptom, error)
	UpdateSymptom(symptom *Symptom) error
	DeleteSymptom(id int64) error
	DeleteAllSymptomsForMeal(mealEntryID int64) error
}
