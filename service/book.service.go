package service

import (
	"Learn-Go-Api/database"
	"Learn-Go-Api/model/entity"
	"time"
)

func GetAllBooks() ([]entity.Book, error) {
	var books []entity.Book
	err := database.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func CreateBook(title, author, cover string, userID uint) (*entity.Book, error) {
	newBook := entity.Book{
		Title:     title,
		Author:    author,
		Cover:     cover,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := database.DB.Create(&newBook).Error; err != nil {
		return nil, err
	}

	return &newBook, nil
}

func GetUserByID(userID uint) (*entity.Users, error) {
	var user entity.Users
	err := database.DB.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
