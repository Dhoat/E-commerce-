package models

import "time"

type Order struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id" gorm:"index"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status" gorm:"type:enum('pending','shipped','delivered');default:'pending'"`
	CreatedAt   time.Time `json:"created_at"`
}
