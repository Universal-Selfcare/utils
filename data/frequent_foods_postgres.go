package data

import (
	"errors"

	"gorm.io/gorm"
)

type PostgresFrequentFoodStore struct {
	DB *gorm.DB
}

func NewPostgresFrequentFoodStore(db *gorm.DB) *PostgresFrequentFoodStore {
	if err := db.AutoMigrate(&FrequentFood{}); err != nil {
		panic("failed to migrate frequent food schema: " + err.Error())
	}
	return &PostgresFrequentFoodStore{DB: db}
}

func (store *PostgresFrequentFoodStore) CreateFrequentFood(food *FrequentFood) (*FrequentFood, error) {
	err := store.DB.Create(food).Error
	if err != nil {
		return nil, err
	}
	return food, nil
}

func (store *PostgresFrequentFoodStore) GetFrequentFood(id int64) (*FrequentFood, error) {
	var food FrequentFood
	err := store.DB.First(&food, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &food, nil
}

func (store *PostgresFrequentFoodStore) ListUserFrequentFoods(userID int64) ([]*FrequentFood, error) {
	var foods []*FrequentFood
	err := store.DB.Where("user_id = ?", userID).Find(&foods).Error
	if err != nil {
		return nil, err
	}
	return foods, nil
}

func (store *PostgresFrequentFoodStore) UpdateFrequentFood(food *FrequentFood) error {
	err := store.DB.Save(food).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresFrequentFoodStore) DeleteFrequentFood(id int64) error {
	err := store.DB.Delete(&FrequentFood{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
