package models

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Total     int       `gorm:"not null" json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Details   []TransactionDetail `gorm:"foreignKey:TransactionID" json:"details,omitempty"`
}
