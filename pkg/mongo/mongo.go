package mongo

import (
	"context"
	"log"
	"ticket-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var Ctx = context.TODO()

func InitMongo() {
	clientOptions := options.Client().ApplyURI(config.Config.MongoUrl)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(Ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database("tickets").Collection("forms")
}