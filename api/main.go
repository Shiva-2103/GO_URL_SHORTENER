package main

import (
	"GO_URL_SHORTENER/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.SendString("This is the API endpoint for POST requests")
	})
	app.Post("/api/v1", routes.ShortenURL) // Existing POST route
}
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New()
	app.Listen(":3000") // Start the app on port 3000

	app.Use(logger.New)

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
