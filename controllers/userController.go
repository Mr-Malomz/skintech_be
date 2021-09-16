package controllers

import (
	"github.com/Mr-Malomz/skintech_be/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = models.GetCollection(models.DB, "user")
var validate = validator.New()
