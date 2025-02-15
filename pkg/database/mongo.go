package database

import (
	"context"
	"fmt"
	"log"

	"banking-ledger/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

// InitMongo initializes MongoDB connection
func InitMongo(cfg *config.Config) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	MongoDB = client.Database("banking_ledger")
	fmt.Println("Connected to MongoDB successfully!")
}
