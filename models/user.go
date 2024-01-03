package models

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/pbkdf2"
)

type User struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Email        string               `bson:"email" json:"email"`
	Hash         string               `bson:"hash" json:"-"`
	Salt         string               `bson:"salt" json:"-"`
	FirstName    string               `bson:"firstName" json:"firstName"`
	LastName     string               `bson:"lastName" json:"lastName"`
	UserLanguage string               `bson:"userLanguage" json:"userLanguage"`
	ImgUrl       string               `bson:"imgUrl" json:"imgUrl"`
	Magic        string               `bson:"magic" json:"magic"`
	Generators   []primitive.ObjectID `bson:"generators" json:"generators"`
	AccountType  string               `bson:"accountType" json:"accountType"`
	Status       string               `bson:"status" json:"status"`
	CreatedAt    time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time            `bson:"updatedAt" json:"updatedAt"`
}

func (u *User) SetPassword(password string) {
	salt := make([]byte, 16)
	_, _ = rand.Read(salt)
	u.Salt = hex.EncodeToString(salt)
	u.Hash = HashPassword(password, u.Salt)
}

func HashPassword(password, salt string) string {
	hash := pbkdf2.Key([]byte(password), []byte(salt), 1000, 64, sha512.New)
	return hex.EncodeToString(hash)
}

func (u *User) GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    u.ID.Hex(),
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 365).Unix(),
		"iat":   time.Now().Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func (u *User) ValidPassword(password string) bool {
	hash := HashPassword(password, u.Salt)
	return u.Hash == hash
}
