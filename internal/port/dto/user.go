package dto

import "time"

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
