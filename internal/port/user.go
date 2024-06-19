package port

import (
	"project/internal/core/domain"
	"time"
)

type UserRepository interface {
	FindUserByKey(string) (*domain.User, error)
	CreateUser(string, string, int) (*domain.User, error)
	CreateKey() string
}

type UserService interface {
	CreateUser(request UserUserRequest) (*UserCreateResponse, error)
}

type UserUserRequest struct {
	Ip       string `json:"ip" binding:"required"`
	ExpireIn int    `json:"expire_in" binding:"required"`
}

type UserCreateResponse struct {
	Url      string    `json:"url"`
	Ip       string    `json:"ip"`
	Key      string    `json:"key"`
	ExpireAt time.Time `json:"expire_at"`
}
