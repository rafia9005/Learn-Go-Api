package handler

import (
	"Learn-Go-Api/database"
	"Learn-Go-Api/model/entity"
	"Learn-Go-Api/model/request"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetBook(c *fiber.Ctx) error {
	var books []entity.Book

	result := database.DB.Find(&books)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	var bookResponses []request.BookResponse
	for _, book := range books {
		bookResponse := request.BookResponse{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
			Cover:  book.Cover,
			UserID: book.UserID,
		}
		bookResponses = append(bookResponses, bookResponse)
	}

	return c.Status(200).JSON(fiber.Map{
		"data": bookResponses,
	})
}

func CreateBook(c *fiber.Ctx) error {
	book := new(request.BookRequest)
	book.Title = c.FormValue("title")
	book.Author = c.FormValue("author")

	userID, err := strconv.ParseUint(c.FormValue("user_id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid user_id",
			"error":   err.Error(),
		})
	}
	book.UserID = uint(userID)

	validate := validator.New()
	if err := validate.Struct(book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	var filePath string
	file, errFile := c.FormFile("cover")
	if errFile == nil {
		filename := file.Filename
		filePath = fmt.Sprintf("/public/images/book/cover/%s", filename)

		errSaveFile := c.SaveFile(file, fmt.Sprintf(".%s", filePath))
		if errSaveFile != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to store file",
				"error":   errSaveFile.Error(),
			})
		}
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filePath,
		UserID: book.UserID,
	}

	if err := database.DB.Create(&newBook).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to store book",
			"error":   err.Error(),
		})
	}

	var user entity.Users
	if err := database.DB.First(&user, newBook.UserID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to fetch user",
			"error":   err.Error(),
		})
	}

	userResponse := request.UsersResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	bookResponse := request.BookResponse{
		ID:        newBook.ID,
		Title:     newBook.Title,
		Author:    newBook.Author,
		Cover:     newBook.Cover,
		User:      userResponse,
		CreatedAt: newBook.CreatedAt,
		UpdatedAt: newBook.UpdatedAt,
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  true,
		"message": "Success creating book",
		"book":    bookResponse,
	})
}
