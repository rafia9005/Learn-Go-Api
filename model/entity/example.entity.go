package entity

import (
	"time"

	"gorm.io/gorm"
)

type Example struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
  Title     string `json:"title"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
