package models

import "time"

type Order struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	BookId    int16     `json:"bookId" gorm:"column:bookId"`
	Point     float64   `json:"point"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
