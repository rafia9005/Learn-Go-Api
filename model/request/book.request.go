package request

type BookResponse struct {
  ID        uint   `gorm:"primaryKey"`
  Title      string `json:"title"`
  Author     string `json:"author"`
  Cover     string `json:"cover"`
}

type BookRequest struct {
	Title      string    `json:"title" validate:"required"`
	Author     string    `json:"author" validate:"required"`
	Cover     string    `json:"cover"`
}
