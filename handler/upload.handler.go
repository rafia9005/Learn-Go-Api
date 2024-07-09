package handler

import (
	"Learn-Go-Api/database"
	"Learn-Go-Api/model/entity"
	"Learn-Go-Api/model/request"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetBook(c *fiber.Ctx) error {
  var book []entity.Book

  result := database.DB.Find(&book)

  if result.Error != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error" : "Internal Server Error",
    })
  }

  var BookResponses []request.BookResponse
  for _, book := range book {
    BookResponse := request.BookResponse{
      ID: book.ID,
      Title: book.Title,
      Author: book.Author,
      Cover: book.Cover,
    }
    BookResponses = append(BookResponses, BookResponse)
  }

  return c.Status(200).JSON(fiber.Map{
    "data" : BookResponses,
  })
}

func CreateBook(c *fiber.Ctx) error {
	book := new(request.BookRequest)

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(book)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   errValidate.Error(),
		})
	}

	// Handle cover
	file, errFile := c.FormFile("cover")

	if errFile != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get file",
			"error":   errFile.Error(),
		})
	}

	filename := file.Filename
	filePath := fmt.Sprintf("/public/images/book/cover/%s", filename)

	errSaveFile := c.SaveFile(file, fmt.Sprintf(".%s", filePath))

	if errSaveFile != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to store file",
			"error":   errSaveFile.Error(),
		})
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filePath,
	}

	errCreate := database.DB.Create(&newBook).Error

	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to store book",
			"error":   errCreate.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  true,
		"message": "Success creating book",
		"book":    newBook,
	})
}
