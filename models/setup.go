package models

import (
	"context"
	"log"
	"time"

	"github.com/Mr-Malomz/skintech_be/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectDB() *mongo.Collection {

	client, err := mongo.NewClient(options.Client().ApplyURI(config.EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("skintech").Collection("skintech")
	
	return collection
}
