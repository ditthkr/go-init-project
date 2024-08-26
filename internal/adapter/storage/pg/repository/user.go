package repository

import (
	"gorm.io/gorm"
	"project/internal/core/domain"
	"project/internal/port"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindUserByKey(key string) (*domain.User, error) {
	user := domain.User{}
	result := r.db.Where("key = ?", key).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
