package routes

import (
	"os"
	"voice_mirror/controllers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupPrivateRoutes(app *fiber.App) {
	// Setup JWT middleware
	var jwtMiddleware = jwtware.New(jwtware.Config{
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		SigningMethod: "HS256",
	})

	app.Post("/private/voice", jwtMiddleware, controllers.CreateVoice)

}
