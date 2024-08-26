package service

import (
	"project/internal/port"
	"project/internal/port/dto"
)

type adminService struct {
	user port.UserRepository
}

func NewUserService(user port.UserRepository) port.UserService {
	return &adminService{user: user}
}

func (r *adminService) CreateUser(req dto.UserUserRequest) (*dto.UserCreateResponse, error) {
	return nil, nil
}
