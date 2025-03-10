package entity

import (
	"time"
)

type User struct {
	ID         uint64
	Username   string
	Nickname   string
	Password   string
	Email      *string
	PictureUrl *string
	Role       uint8
	CreatedAt  time.Time
}
