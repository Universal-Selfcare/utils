package data

import (
	"time"

	"github.com/Universal-Selfcare/utils/validator"
)

var AnonymousUser = &User{}

type User struct {
	ID          int64  `gorm:"primaryKey"                     json:"id"`
	UserName    string `gorm:"type:text;not null;uniqueIndex" json:"user_name"`
	FirstName   string `gorm:"type:text;not null"             json:"first_name"`
	LastName    string `gorm:"type:text;not null"             json:"last_name"`
	Email       string `gorm:"type:text;not null;uniqueIndex" json:"email"`
	PhoneNumber string `gorm:"type:text;not null;uniqueIndex" json:"phone_number"`
	Hash        string `gorm:"type:text;not null"             json:"hash"`

	MedicalInformation MedicalInformation `json:"medical_information"`

	Caregivers         []Caregiver         `json:"caregivers"`
	EmergencyContacts  []EmergencyContact  `json:"emergency_contacts"`
	MedicalEvents      []MedicalEvent      `json:"medical_events"`
	FrequentFoods      []FrequentFood      `json:"frequent_foods"`
	Allergies          []Allergy           `json:"allergies"`
	Medications        []Medication        `json:"medications"`
	DietarySupplements []DietarySupplement `json:"dietary_supplements"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserStore interface {
	CreateUser(user *User) (*User, error)
	GetUser(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByPhoneNumber(phoneNumber string) (*User, error)
	GetByUserName(userName string) (*User, error)
	GetByToken(token string, scope string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int64) error
	ListUsers() ([]*User, error)
}

func (user *User) IsAnonymous() bool {
	return user == AnonymousUser
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.FirstName != "", "first_name", "must be provided")
	v.Check(user.LastName != "", "last_name", "must be provided")
	v.Check(user.PhoneNumber != "", "phone_number", "must be provided")
	v.Check(user.Email != "", "email", "must be provided")
	v.Check(
		validator.Matches(user.PhoneNumber, validator.PhoneRX),
		"phone_number",
		"must be a valid phone number",
	)
	v.Check(
		validator.Matches(user.Email, validator.EmailRX),
		"email",
		"must provide a valid email",
	)
	v.Check(user.UserName != "", "user_name", "must be provided")
	v.Check(
		validator.Matches(user.UserName, validator.UserNameRX),
		"user_name",
		"must be a valid user name (3-16 letters, numbers, underscores, or hyphens)",
	)
}
