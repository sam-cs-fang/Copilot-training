package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	username string
	password float64
}

type UserDto struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password float64 `json:"password"`
}
