package data

import (
	"crypto/sha256"
	"errors"
	"time"

	"gorm.io/gorm"
)

type PostgresTokenStore struct {
	DB *gorm.DB
}

func NewPostgresTokenStore(db *gorm.DB) *PostgresTokenStore {
	if err := db.AutoMigrate(&Token{}); err != nil {
		panic("failed to migrate token schema: " + err.Error())
	}
	return &PostgresTokenStore{DB: db}
}

func (store *PostgresTokenStore) CreateToken(
	userID int64,
	ttl time.Duration,
	scope string,
) (*Token, error) {
	token, err := generateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	err = store.InsertToken(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (store *PostgresTokenStore) InsertToken(token *Token) error {
	err := store.DB.Create(token).Error
	return err
}

func (store *PostgresTokenStore) DeleteAllForUser(scope string, userID int64) error {
	err := store.DB.
		Where("scope = ? AND user_id = ?", scope, userID).
		Delete(&Token{}).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRecordNotFound
	}

	return err
}

func (store *PostgresTokenStore) GetToken(scope string, plaintext string) (*Token, error) {
	hash := sha256.Sum256([]byte(plaintext))

	var token Token

	err := store.DB.
		Where("scope = ? AND hash = ? AND expiry > ?", scope, hash[:], time.Now()).
		First(&token).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &token, nil
}
