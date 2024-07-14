package main

import (
	"Learn-Go-Api/database"
	"Learn-Go-Api/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	// load dotenv
	godotenv.Load()

	// database connect
	database.Connect()

	// auto migration
	routes.AutoMigrate()
	routes.SetupRouter(app)
	port := os.Getenv("APP_PORT")

	if port == "" {
		app.Listen(":3000")
	} else {
		app.Listen(fmt.Sprintf(":%s", port))
	}
}
