package data

import (
	"errors"

	"gorm.io/gorm"
)

type PostgresMedicalEventStore struct {
	DB *gorm.DB
}

func NewPostgresMedicalEventStore(db *gorm.DB) *PostgresMedicalEventStore {
	if err := db.AutoMigrate(&MedicalEvent{}); err != nil {
		panic("failed to migrate medical event schema: " + err.Error())
	}
	return &PostgresMedicalEventStore{DB: db}
}

func (store *PostgresMedicalEventStore) CreateMedicalEvent(
	event *MedicalEvent,
) (*MedicalEvent, error) {
	err := store.DB.Create(event).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (store *PostgresMedicalEventStore) GetMedicalEvent(id int64) (*MedicalEvent, error) {
	var event MedicalEvent
	err := store.DB.First(&event, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (store *PostgresMedicalEventStore) ListUserMedicalEvents(
	userID int64,
) ([]*MedicalEvent, error) {
	var events []*MedicalEvent
	err := store.DB.Where("user_id = ?", userID).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (store *PostgresMedicalEventStore) UpdateMedicalEvent(event *MedicalEvent) error {
	err := store.DB.Save(event).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresMedicalEventStore) DeleteMedicalEvent(id int64) error {
	err := store.DB.Delete(&MedicalEvent{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
