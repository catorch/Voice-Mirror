package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"voice_mirror/models"
)

var validate = validator.New()

func Signup(c *fiber.Ctx) error {
	db := c.Locals("db").(*mongo.Database) // Assuming you've set the DB in the context
	usersCollection := db.Collection("users")

	var data struct {
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required"`
		RPassword string `json:"rpassword" validate:"required,eqfield=Password"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "ERROR", "message": "Cannot parse JSON"})
	}

	if err := validate.Struct(data); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"status": "ERROR", "message": err.Error()})
	}

	var existingUser models.User
	err := usersCollection.FindOne(context.Background(), bson.M{"email": data.Email}).Decode(&existingUser)
	if err == nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"status": "ERROR", "message": "A user with this email already exists"})
	}

	newUser := models.User{
		Email:       data.Email,
		AccountType: "USER",
		Status:      "ACTIVE",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	newUser.SetPassword(data.Password)

	_, err = usersCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "ERROR", "message": "Failed to create user"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "OK", "message": "Account successfully created!"})
}
