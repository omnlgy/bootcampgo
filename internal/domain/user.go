package domain

import (
	"example.com/internal/models"
	"gorm.io/gorm"
)

type User interface{}

type UserRepository interface {
	Create(user User) error
	GetAll() ([]models.User, error)
	GetByID(id uint) (models.User, error)
	GetByEmail(email string) (models.User, error)
	Update(user User) (models.User, error)
	Delete(id uint) error
	WithTx(tx *gorm.DB) UserRepository
}
