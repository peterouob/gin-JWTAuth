package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"jwt-auth/config"
	"log"
	"time"
)

var Client *mongo.Client = DBinstance()

func DBinstance() *mongo.Client {
	conn := config.Config.GetString("mongo.loc")
	if conn == "" {
		conn = "mongodb://localhost:27017"
	}
	ctx, cannel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cannel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		log.Fatalf("Error connecting to mongodb : %v", err)
	}
	return client
}

func OpenCollection(client *mongo.Client, collectioName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("auth").Collection(collectioName)
	return collection
}
