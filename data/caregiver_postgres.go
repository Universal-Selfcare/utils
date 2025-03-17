package data

import (
	"errors"

	"gorm.io/gorm"
)

type PostgresCaregiverStore struct {
	DB *gorm.DB
}

func NewPostgresCaregiverStore(db *gorm.DB) *PostgresCaregiverStore {
	if err := db.AutoMigrate(&Caregiver{}); err != nil {
		panic("failed to migrate caregiver schema: " + err.Error())
	}
	return &PostgresCaregiverStore{DB: db}
}

func (store *PostgresCaregiverStore) CreateCaregiver(caregiver *Caregiver) (*Caregiver, error) {
	err := store.DB.Create(caregiver).Error
	if err != nil {
		return nil, err
	}
	return caregiver, nil
}

func (store *PostgresCaregiverStore) GetCaregiver(id int64) (*Caregiver, error) {
	var caregiver Caregiver
	err := store.DB.First(&caregiver, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &caregiver, nil
}

func (store *PostgresCaregiverStore) GetCaregiverByEmail(email string) (*Caregiver, error) {
	var caregiver Caregiver
	err := store.DB.Where("email = ?", email).First(&caregiver).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &caregiver, nil
}

func (store *PostgresCaregiverStore) ListUserCaregivers(userID int64) ([]*Caregiver, error) {
	var caregivers []*Caregiver
	err := store.DB.Where("user_id = ?", userID).Find(&caregivers).Error
	if err != nil {
		return nil, err
	}
	return caregivers, nil
}

func (store *PostgresCaregiverStore) UpdateCaregiver(caregiver *Caregiver) error {
	err := store.DB.Save(caregiver).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresCaregiverStore) DeleteCaregiver(id int64) error {
	err := store.DB.Delete(&Caregiver{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
