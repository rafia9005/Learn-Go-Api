package handler

import (
	"Learn-Go-Api/model/request"
	"Learn-Go-Api/service"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// GetBooksHandler mengembalikan semua buku.
func GetBooksHandler(c *fiber.Ctx) error {
	books, err := service.GetAllBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch books",
		})
	}

	var bookResponses []request.BookResponse
	for _, book := range books {
		user, err := service.GetUserByID(book.UserID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch user for book",
			})
		}

		bookResponse := request.BookResponse{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
			Cover:  book.Cover,
			User: request.UsersResponse{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
				Role:  user.Role,
			},
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
		}
		bookResponses = append(bookResponses, bookResponse)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": bookResponses,
	})
}

func CreateBookHandler(c *fiber.Ctx) error {
	bookRequest := new(request.BookRequest)

	bookRequest.Title = c.FormValue("title")
	bookRequest.Author = c.FormValue("author")

	userIDStr := c.FormValue("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id",
			"error":   err.Error(),
		})
	}
	bookRequest.UserID = uint(userID)

	validate := validator.New()
	if err := validate.Struct(bookRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	file, err := c.FormFile("cover")
	if err == nil {
		filename := file.Filename
		bookRequest.Cover = fmt.Sprintf("/public/images/book/cover/%s", filename)
		if err := c.SaveFile(file, fmt.Sprintf(".%s", bookRequest.Cover)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to store cover file",
				"error":   err.Error(),
			})
		}
	}

	newBook, err := service.CreateBook(bookRequest.Title, bookRequest.Author, bookRequest.Cover, bookRequest.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create book",
			"error":   err.Error(),
		})
	}

	user, err := service.GetUserByID(bookRequest.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch user",
			"error":   err.Error(),
		})
	}

	bookResponse := request.BookResponse{
		ID:     newBook.ID,
		Title:  newBook.Title,
		Author: newBook.Author,
		Cover:  newBook.Cover,
		User: request.UsersResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
		CreatedAt: newBook.CreatedAt,
		UpdatedAt: newBook.UpdatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Book created successfully",
		"book":    bookResponse,
	})
}
