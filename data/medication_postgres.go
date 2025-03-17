package data

import (
	"errors"

	"gorm.io/gorm"
)

type PostgresMedicationStore struct {
	DB *gorm.DB
}

func NewPostgresMedicationStore(db *gorm.DB) *PostgresMedicationStore {
	if err := db.AutoMigrate(&Medication{}); err != nil {
		panic("failed to migrate medication schema: " + err.Error())
	}
	return &PostgresMedicationStore{DB: db}
}

func (store *PostgresMedicationStore) CreateMedication(medication *Medication) (*Medication, error) {
	err := store.DB.Create(medication).Error
	if err != nil {
		return nil, err
	}
	return medication, nil
}

func (store *PostgresMedicationStore) GetMedication(id int64) (*Medication, error) {
	var medication Medication
	err := store.DB.First(&medication, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &medication, nil
}

func (store *PostgresMedicationStore) ListUserMedications(userID int64) ([]*Medication, error) {
	var medications []*Medication
	err := store.DB.Where("user_id = ?", userID).Find(&medications).Error
	if err != nil {
		return nil, err
	}
	return medications, nil
}

func (store *PostgresMedicationStore) ListUserCurrentMedications(userID int64) ([]*Medication, error) {
	var medications []*Medication
	err := store.DB.Where("user_id = ? AND current = ?", userID, true).Find(&medications).Error
	if err != nil {
		return nil, err
	}
	return medications, nil
}

func (store *PostgresMedicationStore) UpdateMedication(medication *Medication) error {
	err := store.DB.Save(medication).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresMedicationStore) DeleteMedication(id int64) error {
	err := store.DB.Delete(&Medication{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
