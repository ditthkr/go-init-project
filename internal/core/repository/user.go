package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"project/internal/adapter/logs"
	"project/internal/core/domain"
	"project/internal/port"
	"time"
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

func (r *userRepository) CreateUser(key string, ip string, expireIn int) (*domain.User, error) {
	var count int64
	result := r.db.Model(&domain.User{}).Where("ip = ?", ip).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	if count > 0 {
		return nil, fmt.Errorf("user already exists")
	}
	user := domain.User{
		Key:      key,
		Ip:       ip,
		ExpireAt: time.Now().AddDate(0, 0, expireIn),
	}
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateKey() string {
	key := r.randomKey()
	var user *domain.User
	result := r.db.Where("key = ?", key).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return key
		}
		logs.Error(fmt.Sprintf("create key error: %s", result.Error))
		return ""
	}
	return r.CreateKey()
}

func (r *userRepository) randomKey() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.NewSource(time.Now().UnixNano())

	rand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})

	return string(letters[:15])
}
