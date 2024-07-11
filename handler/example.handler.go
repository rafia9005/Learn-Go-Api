package handler

import (
	"Learn-Go-Api/database"
	"Learn-Go-Api/model/entity"
	"Learn-Go-Api/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// get default handler
func GetExample(c *fiber.Ctx) error {
	var example []entity.Example

	result := database.DB.Find(&example)

	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	var ExampleResponses []request.ExampleResponse

	for _, example := range example {
		ExampleResponse := request.ExampleResponse{
			ID:    example.ID,
			Name:  example.Name,
			Title: example.Title,
		}
		ExampleResponses = append(ExampleResponses, ExampleResponse)
	}

	return c.Status(200).JSON(fiber.Map{
		"data": ExampleResponses,
	})
}

// get by id handler
func GetByIdExample(c *fiber.Ctx) error {
	exampleId := c.Params("id")

	var example entity.Example

	result := database.DB.First(&example, exampleId)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Example not found",
		})
	}

	ExampleResponse := request.ExampleResponse{
		ID:    example.ID,
		Name:  example.Name,
		Title: example.Title,
	}

	return c.Status(200).JSON(fiber.Map{
		"data": ExampleResponse,
	})
}

// create handler
func CreateExample(c *fiber.Ctx) error {
	example := new(request.ExampleCreate)

	if err := c.BodyParser(example); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(example)

	if errValidate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed required",
			"error":   errValidate.Error(),
		})
	}

	newExample := entity.Example{
		Name:  example.Name,
		Title: example.Title,
	}

	errCreate := database.DB.Create(&newExample).Error

	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to store Example",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  true,
		"message": "succes create example",
	})
}

// delete handler
func DeleteExample(c *fiber.Ctx) error {
	exampleId := c.Params("id")

	result := database.DB.Delete(&entity.Example{}, exampleId)

	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete Example",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success delete Example",
	})
}

// update handler
func UpdateExample(c *fiber.Ctx) error {
	exampleRequest := new(request.UpdateExample)

	if err := c.BodyParser(exampleRequest); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	exampleId := c.Params("id")
	var example entity.Example
	err := database.DB.First(&example, "id = ?", exampleId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Example not found",
		})
	}

	if exampleRequest.Name != "" {
		example.Name = exampleRequest.Name
	}
	if exampleRequest.Title != "" {
		example.Title = exampleRequest.Title
	}

	responseData := request.ExampleResponse{
		Name:  example.Name,
		Title: example.Title,
	}

	return c.JSON(fiber.Map{
		"message": true,
		"data":    responseData,
	})
}
