package entity

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `json:"title"`
	Author    string `json"author"`
	Cover     string `json:"cover"`
	UserID    uint   `json:"users_id"`
	User      Users  `gorm:foreignKey:UserId`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
