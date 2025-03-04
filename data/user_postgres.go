package data

import (
	"crypto/sha256"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type PostgresUserStore struct {
	DB *gorm.DB
}

func NewPostgresUserStore(db *gorm.DB) *PostgresUserStore {
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate user schema: " + err.Error())
	}
	return &PostgresUserStore{DB: db}
}

func (store *PostgresUserStore) CreateUser(user *User) (*User, error) {
	err := store.DB.Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, ErrRecordConflict
		}
		return nil, err
	}

	return user, nil
}

func (store *PostgresUserStore) GetUser(id int64) (*User, error) {
	var user User
	if err := store.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (store *PostgresUserStore) GetByUserName(userName string) (*User, error) {
	var user User
	if err := store.DB.Where("user_name = ?", userName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (store *PostgresUserStore) GetByEmail(email string) (*User, error) {
	var user User
	if err := store.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (store *PostgresUserStore) GetByPhoneNumber(phoneNumber string) (*User, error) {
	var user User
	if err := store.DB.Where("phone_number = ?", phoneNumber).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (store *PostgresUserStore) UpdateUser(user *User) error {
	if err := store.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (store *PostgresUserStore) DeleteUser(id int64) error {
	if err := store.DB.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (store *PostgresUserStore) ListUsers() ([]*User, error) {
	var users []*User
	if err := store.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (store *PostgresUserStore) GetByToken(scope string, plaintext string) (*User, error) {
	var token Token
	hash := sha256.Sum256([]byte(plaintext))
	err := store.DB.Where("scope = ? AND hash = ? AND expiry > ?", scope, hash[:], time.Now()).
		First(&token).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	var user User
	if err := store.DB.First(&user, token.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &user, nil
}
