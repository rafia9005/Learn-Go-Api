package request

import (
  "time"
)

type ExampleResponse struct {
  	ID        uint   `gorm:"primaryKey"`
  	Name      string `json:"name"`
    Title     string `json:"title"`
}

type ExampleCreate struct {
  Name      string `json:"name" validate:"required"`
  Title     string `json:"title" validate:"required"`
  CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UpdateExample struct {
  Name      string `json:"name" validate:"required"`
  Title     string `json:"title" validate:"required"` 
}


