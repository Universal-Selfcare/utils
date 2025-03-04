package data

import (
	"errors"

	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("record not found")
var ErrRecordConflict= errors.New("new record conflicts with existing record")

type Stores struct {
	UserStore  UserStore
	TokenStore TokenStore
}

func NewStores(db *gorm.DB) *Stores {
	userStore := NewPostgresUserStore(db)
	tokenStore := NewPostgresTokenStore(db)

	return &Stores{
		UserStore:  userStore,
		TokenStore: tokenStore,
	}
}
