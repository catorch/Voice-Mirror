package controllers

import (
	"context"
	"net/http"
	"time"
	"voice_mirror/config"
	"voice_mirror/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateVoice(c *fiber.Ctx) error {
	db := c.Locals("db").(*mongo.Database)
	voicesCollection := db.Collection("voices")

	var data struct {
		Name         string `bson:"email" json:"email" validate:"required"`
		Sample       string `bson:"sample" json:"sample"`
		Accent       string `bson:"accent" json:"accent"`
		Age          string `bson:"age" json:"age"`
		Gender       string `bson:"gender" json:"gender" validate:"required"`
		Language     string `bson:"language" json:"language" validate:"required"`
		LanguageCode string `bson:"languageCode" json:"languageCode" validate:"required"`
		Loudness     string `bson:"loudness" json:"loudness"`
		Style        string `bson:"style" json:"style"`
		Tempo        string `bson:"temp" json:"tempo"`
		Texture      string `bson:"texture" json:"texture"`
		IsCloned     bool   `bson:"isCloned" json:"isCloned" validate:"required"`
		VoiceEngine  string `bson:"voiceEngine" json:"voiceEngine"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "ERROR", "message": "Cannot parse JSON"})
	}

	if err := config.Validator.Struct(data); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"status": "ERROR", "message": err.Error()})
	}

	var existingVoice models.Voice
	err := voicesCollection.FindOne(context.Background(), bson.M{"name": data.Name}).Decode(&existingVoice)
	if err == nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"status": "ERROR", "message": "A voice with this name already exists"})
	}

	newVoice := models.Voice{
		Name:         data.Name,
		Sample:       data.Sample,
		Accent:       data.Accent,
		Age:          data.Age,
		Gender:       data.Gender,
		Language:     data.Language,
		LanguageCode: data.LanguageCode,
		Loudness:     data.Loudness,
		Style:        data.Style,
		Tempo:        data.Tempo,
		IsCloned:     data.IsCloned,
		VoiceEngine:  data.VoiceEngine,
		Status:       "ACTIVE",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err = voicesCollection.InsertOne(context.Background(), newVoice)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "ERROR", "message": "Failed to create user"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "OK", "message": "Voice successfully added"})

}
