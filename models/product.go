package models

import "time"

type Product struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(100);not null" json:"name"`
	Price      int       `gorm:"not null" json:"price"`
	Stock      int       `gorm:"not null;default:0" json:"stock"`
	CategoryID *uint     `gorm:"index" json:"category_id"`
	Category   *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
