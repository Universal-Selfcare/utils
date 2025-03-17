package data

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrRecordConflict = errors.New("new record conflicts with existing record")
)

type Stores struct {
	UserStore               UserStore
	TokenStore              TokenStore
	AllergyStore            AllergyStore
	CaregiverStore          CaregiverStore
	DietarySupplementStore  DietarySupplementStore
	EmergencyContactStore   EmergencyContactStore
	FrequentFoodStore       FrequentFoodStore
	MedicalEventStore       MedicalEventStore
	MedicalInformationStore MedicalInformationStore
	MedicationStore         MedicationStore
}

func NewStores(db *gorm.DB) *Stores {
	userStore := NewPostgresUserStore(db)
	tokenStore := NewPostgresTokenStore(db)
	allergyStore := NewPostgresAllergyStore(db)
	caregiverStore := NewPostgresCaregiverStore(db)
	dietarySupplementStore := NewPostgresDietarySupplementStore(db)
	emergencyContactStore := NewPostgresEmergencyContactStore(db)
	frequentFoodStore := NewPostgresFrequentFoodStore(db)
	medicalEventStore := NewPostgresMedicalEventStore(db)
	medicalInformationStore := NewPostgresMedicalInformationStore(db)
	medicationStore := NewPostgresMedicationStore(db)

	return &Stores{
		UserStore:               userStore,
		TokenStore:              tokenStore,
		AllergyStore:            allergyStore,
		CaregiverStore:          caregiverStore,
		DietarySupplementStore:  dietarySupplementStore,
		EmergencyContactStore:   emergencyContactStore,
		FrequentFoodStore:       frequentFoodStore,
		MedicalEventStore:       medicalEventStore,
		MedicalInformationStore: medicalInformationStore,
		MedicationStore:         medicationStore,
	}
}
