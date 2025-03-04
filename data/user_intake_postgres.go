package data

import (
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


