package test

import (
	"testing"

	"github.com/Mr-Malomz/skintech_be/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestSetupDB(t *testing.T) {
	connect := models.ConnectDB();
	var client *mongo.Client

	if connect != client {
		t.Error()
	}
}
