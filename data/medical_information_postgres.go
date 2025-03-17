package data

import (
	"errors"

	"gorm.io/gorm"
)

type PostgresMedicalInformationStore struct {
	DB *gorm.DB
}

func NewPostgresMedicalInformationStore(db *gorm.DB) *PostgresMedicalInformationStore {
	if err := db.AutoMigrate(&MedicalInformation{}); err != nil {
		panic("failed to migrate medical information schema: " + err.Error())
	}
	return &PostgresMedicalInformationStore{DB: db}
}

func (store *PostgresMedicalInformationStore) CreateMedicalInformation(
	medInfo *MedicalInformation,
) (*MedicalInformation, error) {
	err := store.DB.Create(medInfo).Error
	if err != nil {
		return nil, err
	}
	return medInfo, nil
}

func (store *PostgresMedicalInformationStore) GetMedicalInformation(
	id int64,
) (*MedicalInformation, error) {
	var medInfo MedicalInformation
	err := store.DB.First(&medInfo, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &medInfo, nil
}

func (store *PostgresMedicalInformationStore) GetMedicalInformationByUserID(
	userID int64,
) (*MedicalInformation, error) {
	var medInfo MedicalInformation
	err := store.DB.Where("user_id = ?", userID).First(&medInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &medInfo, nil
}

func (store *PostgresMedicalInformationStore) UpdateMedicalInformation(
	medInfo *MedicalInformation,
) error {
	err := store.DB.Save(medInfo).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresMedicalInformationStore) DeleteMedicalInformation(id int64) error {
	err := store.DB.Delete(&MedicalInformation{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
