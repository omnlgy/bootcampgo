package repository

import (
	"example.com/internal/domain"
	"example.com/internal/models"
	"gorm.io/gorm"
)

var _ domain.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) WithTx(tx *gorm.DB) domain.UserRepository {
	return &UserRepository{db: tx}
}

func (r *UserRepository) Create(user domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	return users, r.db.Find(&users).Error
}

func (r *UserRepository) GetByID(id uint) (models.User, error) {
	var user models.User
	return user, r.db.First(&user, id).Error
}

func (r *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	return user, r.db.Where("email = ?", email).First(&user).Error
}

func (r *UserRepository) Update(user domain.User) (models.User, error) {
	return models.User{}, r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
