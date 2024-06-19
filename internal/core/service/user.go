package service

import (
	"github.com/gofiber/fiber/v2"
	"project/internal/adapter/errs"
	"project/internal/port"
)

type adminService struct {
	user port.UserRepository
}

func NewUserService(user port.UserRepository) port.UserService {
	return &adminService{user: user}
}

func (r *adminService) CreateUser(req port.UserUserRequest) (*port.UserCreateResponse, error) {
	key := r.user.CreateKey()
	user, err := r.user.CreateUser(key, req.Ip, req.ExpireIn)
	if err != nil {
		return nil, errs.Err{Code: fiber.StatusBadRequest, Message: err.Error()}
	}
	return &port.UserCreateResponse{
		Ip:       user.Ip,
		Key:      user.Key,
		ExpireAt: user.ExpireAt,
	}, nil
}
