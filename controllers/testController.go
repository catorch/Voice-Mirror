package controllers

import "github.com/gofiber/fiber/v2"

func Test(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Hello World",
	})
}
