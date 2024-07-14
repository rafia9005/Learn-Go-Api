package providers

import (
	"Learn-Go-Api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	token := c.Get("x-token")

	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	_, err := utils.VerifyToken(token)

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	c.Locals("usersInfo", claims)
	c.Locals("role", claims["role"])

	return c.Next()
}
