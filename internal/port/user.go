package port

import (
	"project/internal/core/domain"
	"project/internal/port/dto"
)

type UserRepository interface {
	FindUserByKey(string) (*domain.User, error)
}

type UserService interface {
	CreateUser(request dto.UserUserRequest) (*dto.UserCreateResponse, error)
}
