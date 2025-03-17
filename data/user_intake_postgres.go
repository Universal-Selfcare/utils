package data

import (
	"errors"

	"gorm.io/gorm"
)

type PostgresUserIntakeStore struct {
	DB *gorm.DB
}

func NewPostgresUserIntakeStore(db *gorm.DB) *PostgresUserIntakeStore {
	if err := db.AutoMigrate(&UserIntake{}); err != nil {
		panic("failed to migrate user intake schema: " + err.Error())
	}
	return &PostgresUserIntakeStore{DB: db}
}

func (store *PostgresUserIntakeStore) CreateUserIntake(
	userIntake *UserIntake,
) (*UserIntake, error) {
	err := store.DB.Create(userIntake).Error
	if err != nil {
		return nil, err
	}
	return userIntake, nil
}

func (store *PostgresUserIntakeStore) GetUserIntakeByUserID(userID int64) (*UserIntake, error) {
	var userIntake UserIntake
	err := store.DB.Where("user_id = ?", userID).First(&userIntake).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &userIntake, nil
}

func (store *PostgresUserIntakeStore) UpdateUserIntake(userIntake *UserIntake) error {
	err := store.DB.Save(userIntake).Error
	if err != nil {
		return err
	}
	return nil
}
