package main

import (
	"context"
	"log"
	"os"
	"voice_mirror/models"
	"voice_mirror/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use(logger.New())

	// Connect to MongoDB
	client := models.ConnectMongoDb()
	defer client.Disconnect(context.Background())

	// Middleware to store the db client in Fiber's local context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", client.Database(os.Getenv("DB_NAME")))
		return c.Next()
	})

	// Setup public routes
	routes.SetupPublicRoutes(app)

	// Start the server
	app.Listen(":3000")
}
