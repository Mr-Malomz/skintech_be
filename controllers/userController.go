package controllers

import (
	"log"

	"github.com/Mr-Malomz/skintech_be/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = models.GetCollection(models.DB, "user")
var validate = validator.New()

func hashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(hashed)
}

func verifyPassword(hashedPassword string, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))

	return err == nil
}

