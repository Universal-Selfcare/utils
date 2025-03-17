package data

import (
	"errors"

	"gorm.io/gorm"
)

// PostgresTrackingPeriodStore implements TrackingPeriodStore interface
type PostgresTrackingPeriodStore struct {
	DB *gorm.DB
}

func NewPostgresTrackingPeriodStore(db *gorm.DB) *PostgresTrackingPeriodStore {
	if err := db.AutoMigrate(&TrackingPeriod{}); err != nil {
		panic("failed to migrate tracking period schema: " + err.Error())
	}
	return &PostgresTrackingPeriodStore{DB: db}
}

func (store *PostgresTrackingPeriodStore) CreateTrackingPeriod(period *TrackingPeriod) (*TrackingPeriod, error) {
	// Check if there's an active tracking period for this user
	var existingPeriod TrackingPeriod
	err := store.DB.Where("user_id = ? AND is_completed = ?", period.UserID, false).First(&existingPeriod).Error
	if err == nil {
		// An active period already exists
		return &existingPeriod, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Some other error occurred
		return nil, err
	}

	// Create a new tracking period
	err = store.DB.Create(period).Error
	if err != nil {
		return nil, err
	}
	return period, nil
}

func (store *PostgresTrackingPeriodStore) GetTrackingPeriod(id int64) (*TrackingPeriod, error) {
	var period TrackingPeriod
	err := store.DB.First(&period, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &period, nil
}

func (store *PostgresTrackingPeriodStore) GetCurrentTrackingPeriod(userID int64) (*TrackingPeriod, error) {
	var period TrackingPeriod
	err := store.DB.Where("user_id = ? AND is_completed = ?", userID, false).First(&period).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &period, nil
}

func (store *PostgresTrackingPeriodStore) GetLastCompletedTrackingPeriod(userID int64) (*TrackingPeriod, error) {
	var period TrackingPeriod
	err := store.DB.Where("user_id = ? AND is_completed = ?", userID, true).
		Order("end_date DESC").
		First(&period).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &period, nil
}

func (store *PostgresTrackingPeriodStore) ListUserTrackingPeriods(userID int64) ([]*TrackingPeriod, error) {
	var periods []*TrackingPeriod
	err := store.DB.Where("user_id = ?", userID).Order("start_date DESC").Find(&periods).Error
	if err != nil {
		return nil, err
	}
	return periods, nil
}

func (store *PostgresTrackingPeriodStore) UpdateTrackingPeriod(period *TrackingPeriod) error {
	err := store.DB.Save(period).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresTrackingPeriodStore) CompleteTrackingPeriod(id int64) error {
	return store.DB.Model(&TrackingPeriod{}).Where("id = ?", id).Update("is_completed", true).Error
}

// PostgresMealEntryStore implements MealEntryStore interface
type PostgresMealEntryStore struct {
	DB *gorm.DB
}

func NewPostgresMealEntryStore(db *gorm.DB) *PostgresMealEntryStore {
	if err := db.AutoMigrate(&MealEntry{}); err != nil {
		panic("failed to migrate meal entry schema: " + err.Error())
	}
	return &PostgresMealEntryStore{DB: db}
}

func (store *PostgresMealEntryStore) CreateMealEntry(entry *MealEntry) (*MealEntry, error) {
	// Check if an entry already exists for this meal
	var existingEntry MealEntry
	err := store.DB.Where(
		"user_id = ? AND tracking_period_id = ? AND tracking_day = ? AND meal_type = ?",
		entry.UserID, entry.TrackingPeriodID, entry.TrackingDay, entry.MealType,
	).First(&existingEntry).Error

	if err == nil {
		// Entry already exists, return it
		return &existingEntry, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Some other error occurred
		return nil, err
	}

	// Create new entry
	err = store.DB.Create(entry).Error
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (store *PostgresMealEntryStore) GetMealEntry(id int64) (*MealEntry, error) {
	var entry MealEntry
	err := store.DB.First(&entry, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (store *PostgresMealEntryStore) GetMealEntryByDetails(
	userID int64,
	trackingPeriodID int64,
	day int,
	mealType string,
) (*MealEntry, error) {
	var entry MealEntry
	err := store.DB.Where(
		"user_id = ? AND tracking_period_id = ? AND tracking_day = ? AND meal_type = ?",
		userID, trackingPeriodID, day, mealType,
	).First(&entry).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (store *PostgresMealEntryStore) ListUserMealEntries(userID int64, trackingPeriodID int64) ([]*MealEntry, error) {
	var entries []*MealEntry
	err := store.DB.Where(
		"user_id = ? AND tracking_period_id = ?",
		userID, trackingPeriodID,
	).Order("tracking_day, meal_type").Find(&entries).Error
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (store *PostgresMealEntryStore) UpdateMealEntry(entry *MealEntry) error {
	err := store.DB.Save(entry).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresMealEntryStore) CompleteMealEntry(id int64) error {
	return store.DB.Model(&MealEntry{}).Where("id = ?", id).Update("is_completed", true).Error
}

func (store *PostgresMealEntryStore) DeleteMealEntry(id int64) error {
	err := store.DB.Delete(&MealEntry{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// PostgresFoodItemStore implements FoodItemStore interface
type PostgresFoodItemStore struct {
	DB *gorm.DB
}

func NewPostgresFoodItemStore(db *gorm.DB) *PostgresFoodItemStore {
	if err := db.AutoMigrate(&FoodItem{}); err != nil {
		panic("failed to migrate food item schema: " + err.Error())
	}
	return &PostgresFoodItemStore{DB: db}
}

func (store *PostgresFoodItemStore) CreateFoodItem(item *FoodItem) (*FoodItem, error) {
	err := store.DB.Create(item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (store *PostgresFoodItemStore) GetFoodItem(id int64) (*FoodItem, error) {
	var item FoodItem
	err := store.DB.First(&item, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (store *PostgresFoodItemStore) GetFoodItemByName(name string) (*FoodItem, error) {
	var item FoodItem
	err := store.DB.Where("name = ?", name).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (store *PostgresFoodItemStore) ListFoodItems() ([]*FoodItem, error) {
	var items []*FoodItem
	err := store.DB.Order("name").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (store *PostgresFoodItemStore) ListFoodItemsByCategory(category string) ([]*FoodItem, error) {
	var items []*FoodItem
	err := store.DB.Where("category = ?", category).Order("name").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (store *PostgresFoodItemStore) UpdateFoodItem(item *FoodItem) error {
	err := store.DB.Save(item).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresFoodItemStore) DeleteFoodItem(id int64) error {
	err := store.DB.Delete(&FoodItem{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// PostgresMealFoodStore implements MealFoodStore interface
type PostgresMealFoodStore struct {
	DB *gorm.DB
}

func NewPostgresMealFoodStore(db *gorm.DB) *PostgresMealFoodStore {
	if err := db.AutoMigrate(&MealFood{}); err != nil {
		panic("failed to migrate meal food schema: " + err.Error())
	}
	return &PostgresMealFoodStore{DB: db}
}

func (store *PostgresMealFoodStore) CreateMealFood(mealFood *MealFood) (*MealFood, error) {
	err := store.DB.Create(mealFood).Error
	if err != nil {
		return nil, err
	}
	return mealFood, nil
}

func (store *PostgresMealFoodStore) GetMealFoodsForMeal(mealEntryID int64) ([]*MealFood, error) {
	var mealFoods []*MealFood
	err := store.DB.Where("meal_entry_id = ?", mealEntryID).Find(&mealFoods).Error
	if err != nil {
		return nil, err
	}
	return mealFoods, nil
}

func (store *PostgresMealFoodStore) DeleteMealFood(id int64) error {
	err := store.DB.Delete(&MealFood{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresMealFoodStore) DeleteAllMealFoodsForMeal(mealEntryID int64) error {
	err := store.DB.Where("meal_entry_id = ?", mealEntryID).Delete(&MealFood{}).Error
	if err != nil {
		return err
	}
	return nil
}

// PostgresCustomFoodStore implements CustomFoodStore interface
type PostgresCustomFoodStore struct {
	DB *gorm.DB
}

func NewPostgresCustomFoodStore(db *gorm.DB) *PostgresCustomFoodStore {
	if err := db.AutoMigrate(&CustomFood{}); err != nil {
		panic("failed to migrate custom food schema: " + err.Error())
	}
	return &PostgresCustomFoodStore{DB: db}
}

func (store *PostgresCustomFoodStore) CreateCustomFood(food *CustomFood) (*CustomFood, error) {
	err := store.DB.Create(food).Error
	if err != nil {
		return nil, err
	}
	return food, nil
}

func (store *PostgresCustomFoodStore) GetCustomFoodsForMeal(mealEntryID int64) ([]*CustomFood, error) {
	var customFoods []*CustomFood
	err := store.DB.Where("meal_entry_id = ?", mealEntryID).Find(&customFoods).Error
	if err != nil {
		return nil, err
	}
	return customFoods, nil
}

func (store *PostgresCustomFoodStore) UpdateCustomFood(food *CustomFood) error {
	err := store.DB.Save(food).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresCustomFoodStore) DeleteCustomFood(id int64) error {
	err := store.DB.Delete(&CustomFood{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresCustomFoodStore) DeleteAllCustomFoodsForMeal(mealEntryID int64) error {
	err := store.DB.Where("meal_entry_id = ?", mealEntryID).Delete(&CustomFood{}).Error
	if err != nil {
		return err
	}
	return nil
}

// PostgresSymptomStore implements SymptomStore interface
type PostgresSymptomStore struct {
	DB *gorm.DB
}

func NewPostgresSymptomStore(db *gorm.DB) *PostgresSymptomStore {
	if err := db.AutoMigrate(&Symptom{}); err != nil {
		panic("failed to migrate symptom schema: " + err.Error())
	}
	return &PostgresSymptomStore{DB: db}
}

func (store *PostgresSymptomStore) CreateSymptom(symptom *Symptom) (*Symptom, error) {
	// Check if a symptom of this type already exists for this meal
	var existingSymptom Symptom
	err := store.DB.Where(
		"meal_entry_id = ? AND symptom_type = ? AND is_overnight = ?",
		symptom.MealEntryID, symptom.SymptomType, symptom.IsOvernight,
	).First(&existingSymptom).Error

	if err == nil {
		// Update existing symptom severity
		existingSymptom.Severity = symptom.Severity
		err = store.DB.Save(&existingSymptom).Error
		if err != nil {
			return nil, err
		}
		return &existingSymptom, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Some other error occurred
		return nil, err
	}

	// Create new symptom
	err = store.DB.Create(symptom).Error
	if err != nil {
		return nil, err
	}
	return symptom, nil
}

func (store *PostgresSymptomStore) GetSymptom(id int64) (*Symptom, error) {
	var symptom Symptom
	err := store.DB.First(&symptom, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &symptom, nil
}

func (store *PostgresSymptomStore) GetSymptomByTypeForMeal(mealEntryID int64, symptomType string) (*Symptom, error) {
	var symptom Symptom
	err := store.DB.Where(
		"meal_entry_id = ? AND symptom_type = ?",
		mealEntryID, symptomType,
	).First(&symptom).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &symptom, nil
}

func (store *PostgresSymptomStore) ListSymptomsForMeal(mealEntryID int64) ([]*Symptom, error) {
	var symptoms []*Symptom
	err := store.DB.Where("meal_entry_id = ?", mealEntryID).Find(&symptoms).Error
	if err != nil {
		return nil, err
	}
	return symptoms, nil
}

func (store *PostgresSymptomStore) UpdateSymptom(symptom *Symptom) error {
	err := store.DB.Save(symptom).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresSymptomStore) DeleteSymptom(id int64) error {
	err := store.DB.Delete(&Symptom{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresSymptomStore) DeleteAllSymptomsForMeal(mealEntryID int64) error {
	err := store.DB.Where("meal_entry_id = ?", mealEntryID).Delete(&Symptom{}).Error
	if err != nil {
		return err
	}
	return nil
}
