package data

import (
	"errors"

	"gorm.io/gorm"
)

type PostgresAllergyStore struct {
	DB *gorm.DB
}

func NewPostgresAllergyStore(db *gorm.DB) *PostgresAllergyStore {
	if err := db.AutoMigrate(&Allergy{}); err != nil {
		panic("failed to migrate allergy schema: " + err.Error())
	}
	return &PostgresAllergyStore{DB: db}
}

func (store *PostgresAllergyStore) CreateAllergy(allergy *Allergy) (*Allergy, error) {
	err := store.DB.Create(allergy).Error
	if err != nil {
		return nil, err
	}
	return allergy, nil
}

func (store *PostgresAllergyStore) GetAllergy(id int64) (*Allergy, error) {
	var allergy Allergy
	err := store.DB.First(&allergy, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &allergy, nil
}

func (store *PostgresAllergyStore) ListUserAllergies(userID int64) ([]*Allergy, error) {
	var allergies []*Allergy
	err := store.DB.Where("user_id = ?", userID).Find(&allergies).Error
	if err != nil {
		return nil, err
	}
	return allergies, nil
}

func (store *PostgresAllergyStore) UpdateAllergy(allergy *Allergy) error {
	err := store.DB.Save(allergy).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresAllergyStore) DeleteAllergy(id int64) error {
	err := store.DB.Delete(&Allergy{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
