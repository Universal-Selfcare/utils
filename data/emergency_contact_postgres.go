package data

import (
	"errors"

	"gorm.io/gorm"
)

type PostgresEmergencyContactStore struct {
	DB *gorm.DB
}

func NewPostgresEmergencyContactStore(db *gorm.DB) *PostgresEmergencyContactStore {
	if err := db.AutoMigrate(&EmergencyContact{}); err != nil {
		panic("failed to migrate emergency contact schema: " + err.Error())
	}
	return &PostgresEmergencyContactStore{DB: db}
}

func (store *PostgresEmergencyContactStore) CreateEmergencyContact(
	contact *EmergencyContact,
) (*EmergencyContact, error) {
	err := store.DB.Create(contact).Error
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (store *PostgresEmergencyContactStore) GetEmergencyContact(
	id int64,
) (*EmergencyContact, error) {
	var contact EmergencyContact
	err := store.DB.First(&contact, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (store *PostgresEmergencyContactStore) GetEmergencyContactByEmail(
	email string,
) (*EmergencyContact, error) {
	var contact EmergencyContact
	err := store.DB.Where("email = ?", email).First(&contact).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (store *PostgresEmergencyContactStore) ListUserEmergencyContacts(
	userID int64,
) ([]*EmergencyContact, error) {
	var contacts []*EmergencyContact
	err := store.DB.Where("user_id = ?", userID).Find(&contacts).Error
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (store *PostgresEmergencyContactStore) UpdateEmergencyContact(
	contact *EmergencyContact,
) error {
	err := store.DB.Save(contact).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresEmergencyContactStore) DeleteEmergencyContact(id int64) error {
	err := store.DB.Delete(&EmergencyContact{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
