package data

import (
	"errors"

	"gorm.io/gorm"
)

type PostgresDietarySupplementStore struct {
	DB *gorm.DB
}

func NewPostgresDietarySupplementStore(db *gorm.DB) *PostgresDietarySupplementStore {
	if err := db.AutoMigrate(&DietarySupplement{}); err != nil {
		panic("failed to migrate dietary supplement schema: " + err.Error())
	}
	return &PostgresDietarySupplementStore{DB: db}
}

func (store *PostgresDietarySupplementStore) CreateDietarySupplement(
	supplement *DietarySupplement,
) (*DietarySupplement, error) {
	err := store.DB.Create(supplement).Error
	if err != nil {
		return nil, err
	}
	return supplement, nil
}

func (store *PostgresDietarySupplementStore) GetDietarySupplement(
	id int64,
) (*DietarySupplement, error) {
	var supplement DietarySupplement
	err := store.DB.First(&supplement, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &supplement, nil
}

func (store *PostgresDietarySupplementStore) ListUserDietarySupplements(
	userID int64,
) ([]*DietarySupplement, error) {
	var supplements []*DietarySupplement
	err := store.DB.Where("user_id = ?", userID).Find(&supplements).Error
	if err != nil {
		return nil, err
	}
	return supplements, nil
}

func (store *PostgresDietarySupplementStore) UpdateDietarySupplement(
	supplement *DietarySupplement,
) error {
	err := store.DB.Save(supplement).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresDietarySupplementStore) DeleteDietarySupplement(id int64) error {
	err := store.DB.Delete(&DietarySupplement{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
