package models

import "time"

type TransactionDetail struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TransactionID uint      `gorm:"not null;index" json:"transaction_id"`
	ProductID     uint      `gorm:"not null;index" json:"product_id"`
	Product       Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity      int       `gorm:"not null" json:"quantity"`
	Subtotal      int       `gorm:"not null" json:"subtotal"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
