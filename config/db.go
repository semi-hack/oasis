package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var RecommendationCollection *mongo.Collection
var UserCollection *mongo.Collection
var CategoryCollection *mongo.Collection

func ConnectDatabase() {
    // Load environment variables from .env file
    if err := godotenv.Load(".env"); err != nil {
        log.Fatalf("Error loading .env file")
    }

    mongoURL := os.Getenv("MONGO_URL")
    if mongoURL == "" {
        log.Fatal("MONGO_URL environment variable not set")
    }

    var err error
    Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Connected to MongoDB!")
    UserCollection = Client.Database("oasis").Collection("users")
    RecommendationCollection = Client.Database("oasis").Collection("recommendation")
	CategoryCollection = Client.Database("oasis").Collection("category")

}