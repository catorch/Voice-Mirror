package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Voice struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"email" json:"email"`
	Sample       string             `bson:"sample" json:"sample"`
	Accent       string             `bson:"accent" json:"accent"`
	Age          string             `bson:"age" json:"age"`
	Gender       string             `bson:"gender" json:"gender"`
	Language     string             `bson:"language" json:"language"`
	LanguageCode string             `bson:"languageCode" json:"languageCode"`
	Loudness     string             `bson:"loudness" json:"loudness"`
	Style        string             `bson:"style" json:"style"`
	Tempo        string             `bson:"temp" json:"tempo"`
	Texture      string             `bson:"texture" json:"texture"`
	IsCloned     bool               `bson:"isCloned" json:"isCloned"`
	VoiceEngine  string             `bson:"voiceEngine" json:"voiceEngine"`
	Status       string             `bson:"isActive" json:"isActive"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
