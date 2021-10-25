package models

import (
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestSetupDB(t *testing.T) {
	connect := ConnectDB()
	var client *mongo.Client

	if connect != client {
		t.Error()
	}
}
