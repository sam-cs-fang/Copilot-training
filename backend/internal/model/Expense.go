package model

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	UserID   int     `gorm:"column:user_id"`
	Title    string  `gorm:"column:title"`
	Amount   float64 `gorm:"column:amount"`
	Category string  `gorm:"column:category"`
	Updater  string  `gorm:"column:updater"`
}

type ExpenseDto struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
}
