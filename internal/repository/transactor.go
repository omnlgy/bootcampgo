package repository

import (
	"gorm.io/gorm"
)

type TransactionRunner interface {
	WithinTransaction(fn func(tx *gorm.DB) error) error
}

// Transactor manages database transactions across multiple repositories.
type Transactor struct {
	db *gorm.DB
}

func NewTransactor(db *gorm.DB) *Transactor {
	return &Transactor{db: db}
}

// WithinTransaction runs fn inside a single DB transaction.
// Commits on success, rolls back on error or panic.
func (t *Transactor) WithinTransaction(fn func(tx *gorm.DB) error) error {
	tx := t.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
