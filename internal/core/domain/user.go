package domain

import "time"

type User struct {
	Id        uint64 `gorm:"primary_key;autoIncrement"`
	Key       string `gorm:"unique;not null"`
	Ip        string `gorm:"not null"`
	Type      string `gorm:"not null;default:'PROXY_ONLY'"` // PROXY_ONLY,PROXY_SCB,SCB,
	Status    bool   `gorm:"not null;default:true"`
	ExpireAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
