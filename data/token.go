package data

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"

	"github.com/Universal-Selfcare/utils/validator"
)

const ScopeAuthentication = "authentication"

type Token struct {
	ID        uint      `json:"-"      gorm:"primary_key"`
	Plaintext string    `json:"token"  gorm:"-"`
	Hash      []byte    `json:"-"      gorm:"not null;uniqueIndex"`
	UserID    int64     `json:"-"      gorm:"not null;index"`
	Expiry    time.Time `json:"expiry" gorm:"not null;index"`
	Scope     string    `json:"-"      gorm:"not null;index"`
}

type TokenStore interface {
	CreateToken(userID int64, ttl time.Duration, scope string) (*Token, error)
	InsertToken(token *Token) error
	DeleteAllForUser(scope string, userID int64) error
	GetToken(scope string, plaintext string) (*Token, error)
}

func generateToken(userID int64, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]

	return token, nil
}

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long")
}
