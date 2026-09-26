package repository

import (
	"errors"

	"github.com/vitaly06/geo-weather-bot/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id int64) (*domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "id = ?", id).Error

	return &user, err
}

func (r *UserRepository) Exists(id int64) (bool, error) {
	var count int64

	err := r.db.Model(&domain.User{}).Where("id = ?", id).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
