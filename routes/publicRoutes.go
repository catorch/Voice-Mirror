package routes

import (
	"voice_mirror/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPublicRoutes(app *fiber.App) {

	app.Post("/public/auth/signup", controllers.Signup)

	app.Get("/test", controllers.Test)
}
